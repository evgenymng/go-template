import threading
import time
import random
import requests
import signal
import os
from datetime import datetime

# configuration
BASE_URL = os.environ.get("BASE_URL", "http://localhost:8080")
METRICS_ENDPOINT = f"{BASE_URL}/send-metrics"
TRACE_ENDPOINT = f"{BASE_URL}/send-trace"
LOGS_ENDPOINT = f"{BASE_URL}/send-logs"

# global flag to stop threads
running = True


def signal_handler(sig, frame):
    global running
    print("\nShutting down gracefully...")
    running = False


def log_request(endpoint, response_time, status_code, thread_name):
    timestamp = datetime.now().strftime("%H:%M:%S")
    print(
        f"[{timestamp}] {thread_name}: {endpoint} -> {status_code} ({response_time:.2f}s)"
    )


def metrics_worker():
    """Worker thread for /send-metrics endpoint"""
    thread_name = "METRICS"
    request_count = 0

    while running:
        try:
            start_time = time.time()
            response = requests.get(METRICS_ENDPOINT, timeout=5)
            response_time = time.time() - start_time

            request_count += 1
            log_request(
                "/send-metrics", response_time, response.status_code, thread_name
            )

            if response.status_code == 200:
                # Optional: log response data
                data = response.json()
                print(f"    └── Metrics: {data.get('metrics', {})}")

        except requests.exceptions.RequestException as e:
            print(f"[{thread_name}] Error: {e}")
        except Exception as e:
            print(f"[{thread_name}] Unexpected error: {e}")

        # Random sleep between 0.5 and 2.0 seconds
        if running:
            sleep_time = random.uniform(0.5, 2.0)
            time.sleep(sleep_time)

    print(f"[{thread_name}] Stopped after {request_count} requests")


def trace_worker():
    """Worker thread for /send-trace endpoint"""
    thread_name = "TRACE"
    request_count = 0

    while running:
        try:
            start_time = time.time()
            response = requests.get(TRACE_ENDPOINT, timeout=5)
            response_time = time.time() - start_time

            request_count += 1
            log_request("/send-trace", response_time, response.status_code, thread_name)

            if response.status_code == 200:
                # Optional: log response data
                data = response.json()
                trace_info = data.get("trace_info", {})
                print(f"    └── Trace ID: {trace_info.get('trace_id', 'N/A')[:16]}...")

        except requests.exceptions.RequestException as e:
            print(f"[{thread_name}] Error: {e}")
        except Exception as e:
            print(f"[{thread_name}] Unexpected error: {e}")

        # Random sleep between 0.5 and 2.0 seconds
        if running:
            sleep_time = random.uniform(0.5, 2.0)
            time.sleep(sleep_time)

    print(f"[{thread_name}] Stopped after {request_count} requests")


def logs_worker():
    """Worker thread for /send-logs endpoint"""
    thread_name = "LOGS"
    request_count = 0

    while running:
        try:
            start_time = time.time()
            response = requests.get(LOGS_ENDPOINT, timeout=5)
            response_time = time.time() - start_time

            request_count += 1
            log_request("/send-logs", response_time, response.status_code, thread_name)

            if response.status_code == 200:
                # Optional: log response data
                data = response.json()
                trace_info = data.get("trace_info", {})
                logging_demo = data.get("logging_demo", {})
                print(f"    └── Trace ID: {trace_info.get('trace_id', 'N/A')[:16]}... | Logs: {logging_demo.get('context_aware_logs', 0)}")

        except requests.exceptions.RequestException as e:
            print(f"[{thread_name}] Error: {e}")
        except Exception as e:
            print(f"[{thread_name}] Unexpected error: {e}")

        # Random sleep between 0.5 and 2.0 seconds
        if running:
            sleep_time = random.uniform(0.5, 2.0)
            time.sleep(sleep_time)

    print(f"[{thread_name}] Stopped after {request_count} requests")


def main():
    global running

    # Set up signal handler for graceful shutdown
    signal.signal(signal.SIGINT, signal_handler)

    print("Starting OpenTelemetry demo load tester...")
    print(f"Metrics endpoint: {METRICS_ENDPOINT}")
    print(f"Trace endpoint: {TRACE_ENDPOINT}")
    print(f"Logs endpoint: {LOGS_ENDPOINT}")
    print("Press Ctrl+C to stop")
    print("-" * 50)

    # Create and start threads
    metrics_thread = threading.Thread(target=metrics_worker, name="MetricsWorker")
    trace_thread = threading.Thread(target=trace_worker, name="TraceWorker")
    logs_thread = threading.Thread(target=logs_worker, name="LogsWorker")

    metrics_thread.start()
    trace_thread.start()
    logs_thread.start()

    try:
        # Wait for threads to complete
        metrics_thread.join()
        trace_thread.join()
        logs_thread.join()
    except KeyboardInterrupt:
        running = False
        print("\nWaiting for threads to finish...")
        metrics_thread.join()
        trace_thread.join()
        logs_thread.join()

    print("Load tester stopped.")


if __name__ == "__main__":
    main()
