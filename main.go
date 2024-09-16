package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type InvoiceData struct {
	MonthRange               string
	CurrentDateFormatted     string
	DaysInMonth              int
	CurrentDateTimeFormatted string
	InvoiceNumber            string
	Amount                   string
	EmailTitle               string
	EmailBody                string
	AddressLine1             string
	AddressLine2             string
	City                     string
	PostalCode               string
	Country                  string
	Name                     string
	BillTo                   string
	BillToAddress            string
	Bank                     string
	BankBranch               string
	Account_No               string
	SwiftCode                string
	BankCode                 string
	BranchCode               string
}

// GeneratePDF creates a PDF from an HTML template and data.
func generatePDF(data InvoiceData) error {
	// Load and parse the HTML template
	tmpl, err := template.ParseFiles("invoice.html")
	if err != nil {
		return err
	}

	// Create a buffer to hold the rendered HTML
	var renderedHTML bytes.Buffer
	err = tmpl.Execute(&renderedHTML, data)
	if err != nil {
		return err
	}

	// Convert HTML to PDF
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(&renderedHTML))
	err = pdfg.Create()
	if err != nil {
		return err
	}

	// Save the PDF to a file
	err = pdfg.WriteFile(data.InvoiceNumber + ".pdf")
	if err != nil {
		return err
	}

	return nil
}

func sendEmail(data InvoiceData) error {
	// Email configurations
	from := os.Getenv("SMTP_USER")
	to := os.Getenv("RECIPIENTS")
	subject := data.EmailTitle
	body := data.EmailBody

	// Gmail SMTP settings
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := 587 // Use the port as an integer
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	// Split the recipients into a slice
	recipients := strings.Split(to, ",")

	// Check if the PDF file exists before attaching it
	if _, err := os.Stat(data.InvoiceNumber + ".pdf"); os.IsNotExist(err) {
		return fmt.Errorf("PDF file not found: %v", err)
	}

	// Create a new email message
	m := gomail.NewMessage()
	if m == nil {
		return fmt.Errorf("failed to initialize email message")
	}

	m.SetHeader("From", from)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Add multiple recipients
	m.SetHeader("To", recipients...)

	// Attach the generated PDF
	m.Attach(data.InvoiceNumber + ".pdf")

	// Create a new SMTP client and check for errors during initialization
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	if d == nil {
		return fmt.Errorf("failed to create SMTP dialer")
	}

	// Send the email and check for errors
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error while sending email: %v", err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}

func getPayload() InvoiceData {
	// Get current time
	currentTime := time.Now()

	// Get the first and last day of the current month
	firstDay := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.Local)
	lastDay := time.Date(currentTime.Year(), currentTime.Month()+1, 1, 0, 0, 0, 0, time.Local).Add(-time.Second)

	month := currentTime.Format("January")
	// Assign values to the struct fields
	dateTimeInfo := InvoiceData{
		MonthRange:               fmt.Sprintf("%s - %s", firstDay.Format("January 2, 2006"), lastDay.Format("January 2, 2006")),
		CurrentDateFormatted:     currentTime.Format("02.01.2006"),
		DaysInMonth:              lastDay.Day(),
		CurrentDateTimeFormatted: currentTime.Format("1/2/2006 15:04:05"),
		InvoiceNumber:            os.Getenv("EMPLOYEE_NUMBER") + currentTime.Format("200601"),
		Amount:                   "$" + os.Getenv("AMOUNT"),
		EmailTitle:               month + " Invoice - Dilshan",
		EmailBody:                "Dear Accounts Team,\n\nPlease refer to the attached invoice for " + month + ". \n\nThank you",
		AddressLine1:             os.Getenv("ADDRESS_LINE_1"),
		AddressLine2:             os.Getenv("ADDRESS_LINE_2"),
		City:                     os.Getenv("CITY"),
		PostalCode:               os.Getenv("POSTAL_CODE"),
		Country:                  os.Getenv("COUNTRY"),
		Name:                     os.Getenv("NAME"),
		BillTo:                   os.Getenv("BILL_TO"),
		BillToAddress:            os.Getenv("BILL_TO_ADDRESS"),
		Bank:                     os.Getenv("BANK"),
		BankBranch:               os.Getenv("BANK_BRANCH"),
		Account_No:               os.Getenv("ACCOUNT_NO"),
		SwiftCode:                os.Getenv("SWIFT_CODE"),
		BankCode:                 os.Getenv("BANK_CODE"),
		BranchCode:               os.Getenv("BRANCH_CODE"),
	}

	return dateTimeInfo
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file: %v", err)
	}

	data := getPayload()

	// Step 1: Generate the PDF
	err = generatePDF(data)
	if err != nil {
		fmt.Println("Failed to generate PDF:", err)
		os.Exit(1)
	}

	// Step 2: Send the email
	err = sendEmail(data)
	if err != nil {
		fmt.Println("Failed to send email:", err)
		os.Exit(1)
	}
}
