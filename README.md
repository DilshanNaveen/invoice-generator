# Invoice Generator

This project is a Go-based tool that generates invoices in PDF format and sends them via email. It uses an HTML template to create a styled invoice PDF and sends the generated invoice using SMTP settings provided via environment variables.

## Features

- Generate invoices from a customizable HTML template.
- Dynamically load invoice data from environment variables.
- Convert HTML to PDF using `wkhtmltopdf`.
- Send the invoice via email using the SMTP protocol.

## Technologies Used

- **Go (Golang)**: Main programming language for the project.
- **wkhtmltopdf**: A tool to convert HTML to PDF.
- **gomail.v2**: A simple library to send emails from Go.
- **SebastiaanKlippert/go-wkhtmltopdf**: A Go wrapper for `wkhtmltopdf`.
- **HTML Templates**: For invoice formatting.

## Prerequisites

Before running the project, ensure you have the following:

1. **Go**: [Download and install Go](https://golang.org/doc/install) if you haven't already.
2. **wkhtmltopdf**: [Install `wkhtmltopdf`](https://wkhtmltopdf.org/downloads.html), as it's required for generating PDFs.

### Installation Instructions for `wkhtmltopdf`

- **macOS**:
  ```bash
  brew install wkhtmltopdf
  ```

- **Ubuntu**:
  ```bash
  sudo apt-get update
  sudo apt-get install -y wkhtmltopdf
  ```

- **Windows**:
  Download and install `wkhtmltopdf` from [here](https://wkhtmltopdf.org/downloads.html).

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/DilshanNaveen/invoice-generator.git
cd invoice-generator
```

### 2. Install Dependencies

Make sure to initialize the Go module and install necessary dependencies:

```bash
go mod tidy
```

### 3. Set Up Environment Variables

You need to configure environment variables to load essential invoice and email data. You can do this by creating a `.env` file in the project root with the following variables:

```bash
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASS=your-email-password
RECIPIENTS=recipient@example.com
EMPLOYEE_NUMBER=12345
AMOUNT=1000
BANK=Your Bank
BANK_BRANCH=Main Branch
ACCOUNT_NO=9876543210
SWIFT_CODE=ABCDEFGH
BANK_CODE=123
BRANCH_CODE=456
NAME=Your Name
ADDRESS_LINE_1=Street 1
ADDRESS_LINE_2=Street 2
CITY=Your City
POSTAL_CODE=12345
COUNTRY=Your Country
BILL_TO=Client Name
BILL_TO_ADDRESS=Client Address
```

Alternatively, you can load these variables manually in your shell or pass them when running the Go script (see [Running Locally](#running-locally)).

### 4. Running Locally

After setting up your environment variables, you can run the Go project with the following command:

```bash
go run main.go
```

Make sure `wkhtmltopdf` is installed and available in your system's path.

## Usage

1. **Generating and Sending an Invoice**: The program will automatically:
   - Generate a PDF invoice using the provided HTML template and data.
   - Email the invoice to the recipient(s) specified in the environment variables.

2. **Customize the Invoice Template**:
   The HTML template used to generate the invoice is `invoice.html`. Modify this file according to your design preferences.

## Environment Variables

| Variable          | Description                                    |
| ----------------- | ---------------------------------------------- |
| `SMTP_HOST`       | The SMTP server address (e.g., `smtp.gmail.com`). |
| `SMTP_PORT`       | SMTP port (commonly `587` for Gmail).           |
| `SMTP_USER`       | Your SMTP username (usually your email).        |
| `SMTP_PASS`       | Your SMTP password.                             |
| `RECIPIENTS`      | Comma-separated list of recipient email addresses. |
| `EMPLOYEE_NUMBER` | Unique identifier for generating the invoice number. |
| `AMOUNT`          | Amount to be billed.                            |
| `BANK`            | The name of your bank.                          |
| `BANK_BRANCH`     | Bank branch details.                            |
| `ACCOUNT_NO`      | Your account number.                            |
| `SWIFT_CODE`      | SWIFT code for international transfers.         |
| `BANK_CODE`       | Bank code for identification.                   |
| `BRANCH_CODE`     | Branch code for the bank branch.                |
| `NAME`            | Your full name or company name.                 |
| `ADDRESS_LINE_1`  | Address line 1 (e.g., street address).          |
| `ADDRESS_LINE_2`  | Address line 2 (optional).                      |
| `CITY`            | City name.                                      |
| `POSTAL_CODE`     | Postal/ZIP code.                                |
| `COUNTRY`         | Country of residence.                           |
| `BILL_TO`         | The client or entity being billed.              |
| `BILL_TO_ADDRESS` | Address of the client or entity.                |

## Project Structure

```bash
.
├── main.go                 # Main Go program file
├── go.mod                  # Go module file
├── invoice.html            # HTML template for invoice
├── README.md               # Project documentation
└── .env                    # Environment variables (not included in version control)
```

## Deployment

This project includes a GitHub Actions workflow that automates the PDF generation and email sending process when changes are pushed to the `main` branch.

### GitHub Actions Workflow

The CI/CD pipeline is set up to:

1. Install dependencies (`go mod tidy`).
2. Install `wkhtmltopdf` on the runner.
3. Run the Go script with the environment variables passed via GitHub Secrets.

For more information, check the `.github/workflows` folder in the project repository.

## Contributing

If you'd like to contribute to this project, feel free to submit a pull request. Any feedback and improvements are welcome!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.