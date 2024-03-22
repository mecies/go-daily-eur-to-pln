# Daily EUR to PLN 🤑

This project is a cron job written in Go that runs every day at 2 PM UTC. It checks the exchange rate of EUR to PLN and sends an email notification to a specified receiver.

## Features 🚀

- Runs as a cron job to check EUR to PLN exchange rate daily.
- Sends email notification with the exchange rate to a specified receiver.

## Setup ⚙️

1. Clone the repository:

   ```bash
   git clone https://github.com/mecies/go-daily-eur-to-pln.git
   ```

2. Build the project:

   ```bash
   go build
   ```

3. Set up environment variables for email credentials:

   ```bash
    EMAIL_PASSWORD=your-email-password
    EMAIL_SENDER=your-sender-email@something.com
    EMAIL_RECEIVER=your-receiver-email@something.com
    EMAIL_HOST=smtp.your.host
    EMAIL_PORT=587
   ```

4. Run the project:

   ```bash
   ./go-daily-eur-to-pln
   ```

   
## License 📜

This project is licensed under the [MIT License](LICENSE).
