# <u>Simple Email Checker Using Golang</u>

This is a simple project built using the built-in net/http package in golang that can be widely used to deliver email security, prevent spam, authentical emails and prevent phishing attacks as well.

The project uses the built-in functions in the net/http package in golang that are described below with their corresponding use cases: 

* net.LookupMX(): This looks up if the email has a MX (Mail Exchanger) record, which specifies the mail servers responsible for receiving emails for your domain. The absence of a MX record signifies an inauthentic domain.

* net.LookupTXT(): It helps look for email servers that are allowed to send emails on behalf of that given domain, if the lookup contains a prefix spf1. If the spf test fails, the server looks up the dmarc record to determine what to do with the server that failed the spf test. 

![Go_Image](https://gitconnected.com/public/images/tutorials/golang)

Thank You for visiting!
