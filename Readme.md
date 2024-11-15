# Burger Service Appointment Parser

This Go program retrieves and displays the next available appointment date for various municipal services in Bremen, Germany. It allows users to specify the service URL as a command-line argument, making it adaptable for multiple services. Additionally, it supports integration with notification systems and periodic scheduling.

## Features

- Fetches the earliest available appointment date from a user-provided URL.
- Parses and extracts the appointment date from the webpage content.
- Displays the next available appointment date in a user-friendly format.
- Extensible to handle different municipal services beyond passport applications.
- Output can be piped into the [Telegram Notifier](https://github.com/phpanhey/telegram_notify) project for instant notifications.
- Can be run periodically via CronTab to ensure you never miss an appointment.

## Prerequisites

- [Go](https://golang.org/dl/) installed on your system.
- Internet connection to access the service website.

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/phpanhey/burgerserviceappointmentparster.git
   ```

2. Navigate to the project directory:

   ```bash
   cd burgerserviceappointmentparster
   ```

3. Build the program. The `go.mod` file will automatically fetch all dependencies:

   ```bash
   go build
   ```

## Usage

### Running the Program

Execute the program with the desired service URL as a command-line argument:

```bash
./burgerserviceappointmentparster <URL>
```

Replace `<URL>` with the actual URL of the service page you want to check. The program will output:

```
Next available appointment date: DD.MM.YYYY
```

### Piping Output to Telegram Notifier

You can pipe the output of this program into your [Telegram Notifier](https://github.com/phpanhey/telegram_notify) project to receive instant updates on new appointments. For example:

```bash
./burgerserviceappointmentparster <URL> | python3 telegram_notify.py --message "$(cat -)" --telegram_bot_token "<your_bot_token>" --telegram_user_id "<your_user_id>"
```

This integration allows you to stay updated on appointment availability in real time.

### Scheduled Runs with CronTab

To ensure you never miss an appointment, you can schedule this program to run periodically using CronTab. Here's an example of a CronTab entry to run the script every hour:

```bash
0 * * * * /path/to/burgerserviceappointmentparster <URL> | python3 /path/to/telegram_notify.py --message "$(cat -)" --telegram_bot_token "<your_bot_token>" --telegram_user_id "<your_user_id>"
```

Replace `/path/to/` with the actual paths to the respective scripts and `<URL>` with the service URL.

## How It Works

1. Fetches the content of the specified URL.
2. Searches for the section containing the earliest available appointment date.
3. Parses the HTML to extract and clean the date.
4. Outputs the date in a readable format.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [soup](https://github.com/anaskhan96/soup): A Go library for web scraping.

*Note: This program is intended for informational purposes only. Appointment availability may change, and it is recommended to verify the information through official channels.*
```