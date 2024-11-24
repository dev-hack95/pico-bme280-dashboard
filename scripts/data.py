import serial
import sqlite3
import signal
import sys
from datetime import datetime
import logging
from contextlib import closing

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s'
)

class SensorDataLogger:
    def __init__(self, port: str = "/dev/ttyACM0", baudrate: int = 115200, db_path: str = '../records.db'):
        self.port = port
        self.baudrate = baudrate
        self.db_path = db_path
        self.running = True
        
        signal.signal(signal.SIGINT, self.handle_shutdown)
        signal.signal(signal.SIGTERM, self.handle_shutdown)
        
    def handle_shutdown(self, signum, frame):
        logging.info("Shutdown signal received. Closing connections...")
        self.running = False
        
    def setup_database(self) -> sqlite3.Connection:
        """Initialize database connection and create table if needed."""
        conn = sqlite3.connect(self.db_path)
        with closing(conn.cursor()) as cursor:
            cursor.execute('''
                CREATE TABLE IF NOT EXISTS sensor_readings (
                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                    humidity TEXT,
                    pressure TEXT,
                    temperature TEXT,
                    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
                    active BOOLEAN DEFAULT TRUE
                )
            ''')
        conn.commit()
        return conn

    def parse_sensor_data(self, line: str):
        try:
            parts = [part.strip() for part in line.split(',')]
            if len(parts) != 3:
                logging.warning(f"Invalid data format: {line}")
                return None
                
            humidity = parts[0].rstrip('%')
            pressure = parts[1]
            temperature = parts[2]
            
            return humidity, pressure, temperature
            
        except Exception as err:
            logging.error(f"Error parsing line '{line}': {err}")
            return None

    def process_data(self):
        try:
            with closing(serial.Serial(self.port, self.baudrate)) as serial_conn, \
                 closing(self.setup_database()) as db_conn, \
                 closing(db_conn.cursor()) as cursor:
                
                buffer = ""
                logging.info("Starting data collection. Press Ctrl+C to stop...")
                
                while self.running:
                    try:
                        if serial_conn.in_waiting:
                            chunk = serial_conn.read(serial_conn.in_waiting).decode('utf-8')
                            buffer += chunk
                            
                            while '\n' in buffer:
                                line, buffer = buffer.split('\n', 1)
                                if line.strip():
                                    sensor_data = self.parse_sensor_data(line)
                                    if sensor_data:
                                        humidity, pressure, temperature = sensor_data
                                        cursor.execute('''
                                            INSERT INTO sensor_readings (humidity, pressure, temperature, active)
                                            VALUES (?, ?, ?, ?)
                                        ''', (humidity, pressure, temperature, True))
                                        db_conn.commit()
                                        logging.info(f"Stored reading: {temperature}, {humidity}%, {pressure}")
                                        
                    except serial.SerialException as err:
                        logging.error(f"Serial port error: {err}")
                        self.running = False
                        
        except sqlite3.Error as err:
            logging.error(f"Database error: {err}")
        except Exception as err:
            logging.error(f"Unexpected error: {err}")
        finally:
            logging.info("Data collection stopped.")

def main():
    try:
        logger = SensorDataLogger()
        logger.process_data()
    except Exception as err:
        logging.error(f"Fatal error: {err}")
    finally:
        logging.info("Program terminated.")
        sys.exit(0)

if __name__ == "__main__":
    main()
