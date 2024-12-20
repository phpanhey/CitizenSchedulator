# CitizenSchedulator

CitizenSchedulator is a Go-based application designed to help users find the next available appointment for various municipal services in Bremen, Germany. It fetches the earliest appointment date from a specified service URL and can integrate with notification systems for real-time updates.

## Features

- Retrieves and displays the next available appointment date from a user-provided URL.
- Supports multiple municipal services.
- Extensible for various use cases beyond passport applications.
- Outputs results in a user-friendly format.
- Integrates seamlessly with the [Telegram Notifier](https://github.com/phpanhey/telegram_notify) for instant notifications.
- Can be scheduled to run periodically using CronTab.

## Prerequisites

- [Go](https://golang.org/dl/) installed on your system.
- Internet connection to access the service website.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/phpanhey/citizenschedulator.git
   ```
2. Navigate to the project directory:
   ```bash
   cd citizenschedulator
   ```
3. Build the program. Dependencies will be fetched automatically using the `go.mod` file:
   ```bash
   go build
   ```

## Usage

### Running the Program
Run the application with a service URL as a command-line argument:
```bash
./citizenschedulator <service_url>
```
Example output:
```
Next available appointment date: DD.MM.YYYY
```

### Piping Output to Telegram Notifier
Integrate with [Telegram Notifier](https://github.com/phpanhey/telegram_notify) to receive notifications:
```bash
python3 telegram_notify.py --message "$(./citizenschedulator <service_url>)" --telegram_bot_token "<your_bot_token>" --telegram_user_id "<your_user_id>"
```

### Scheduled Runs with CronTab
Automate periodic checks using CronTab. For example, to run the script hourly:
```bash
0 * * * * cd /path/to/CitizenSchedulator; python3 /path/to/telegram_notify.py --message "$(./citizenschedulator <service_url>)" --telegram_bot_token "<your_bot_token>" --telegram_user_id "<your_user_id>"; cd;
```
Replace `/path/to/` with the actual script paths and `<service_url>` with the service URL.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [soup](https://github.com/anaskhan96/soup): A Go library for web scraping.