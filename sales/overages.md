:orphan:
:nosearch:

# Gravwell Data Ingestion Limits and Overages Policy
Last Updated: June, 2025

This Data Ingestion Limits and Overages Policy (the **“Overages Policy”**) supplements and is incorporated by reference into the Gravwell Software and Services Agreement available at [https://docs.gravwell.io/eula.html](https://docs.gravwell.io/eula.html) (the **“Agreement”**) entered into between Gravwell and the organization and/or or entity specified on the applicable Order Form.  Capitalized terms not otherwise defined herein have the meanings set forth in the Agreement.

## 1. Daily Data Ingestion Limits:
During Customer’s Subscription Term, Customer is permitted to ingest the volumes of data specified in the applicable Order Form into the Software measured over a 24-hour calendar day (the **“Daily Ingestion Limit”**).  Reingestion of data shall count towards the Daily Ingestion Limit.
Customer may request an increase in the Daily Ingestion Limit at any time, provided that, such increase shall continue for the remainder of Customer’s then-current Subscription Term and subject to Customer’s payment of additional license fees for such increase, prorated for the remainder of Customer’s then-current Subscription Term.  In addition, Customer may request a temporary increase in their Daily Ingestion Limit (**"Temporary Capacity Increase"**) for specific one-off projects, subject to and in accordance with Section 5. Any increase in the Daily Ingestion Limits, whether a permanent increase for the remainder of the Subscription Term or a Temporary Capacity Increase (each, an **“Increase”**) and the applicable fees for such Increase, shall be mutually agreed upon in writing (including, via an Order Form).
The Daily Ingestion Limits and Temporary Capacity Limits (as defined in Section 5.2) are deemed Usage Limitations as such term is defined and used in the Agreement.

## 2. Retention Period:
Customer’s data is subject to a retention period specified in the applicable Order Form (the **“Retention Period”**). After the Retention Period, such data may be automatically archived or deleted in accordance with Gravwell’s then-current data retention policy.  Customer will provide an S3-compatible storage endpoint for data archival.  Data will be archived in an open format with all tools necessary to reingest and analyze.

## 3. Monitoring and Reporting:
Customer will have access to Customer’s daily ingestion volume via a reporting tool through the user interface of the Software for Customer to monitor and ensure its data ingestion complies with the Daily Ingestion Limit.  The customer is responsible for monitoring its daily ingestion volumes regularly.

## 4. Overages:
### 4.1. _General._  
If Customer's daily data ingestion exceeds the Daily Ingestion Limit on any given day (**"Overage"**), Gravwell will use reasonable efforts to notify Customer via email within 48 hours following such Overage.  

### 4.2. _Excessive Overages._  
Overages exceeding the limits in the chart below (**“Excessive Overages”**) will be handled in accordance with the following process:
| Phase    | Excessive Overages       | Action              |
|----------|:------------------------:|---------------------|
| Phase 1  |<p>> 15% Overage within any 14-day period<br/>**OR**<br/>any Overages for more than 7 days</p> |<p>Customer’s designated Gravwell customer solutions architect will notify Customer via email of the Excessive Overage (the **“Initial Notice”**) and share a report breaking down the data and cause of the Overage and will arrange a call with Customer to understand the causes of the Overage and whether the Customer elects to either: (1) reduce its usage to comply with the Daily Ingestion Limits (a **“Reduction”**), or (2) Increase its Daily Ingestion Limits.<br/>If Customer elects an Increase, then any Overages following the Increase shall be considered new Overages and will be addressed starting at Phase 1.</p> |
| Phase 2  |<p>***Following Customer’s election of a Reduction or Customer’s failure to notify Gravwell of its election within 10 days following the Initial Notice:***<br/>> 15% Overage within any 30-day period<br/>**OR**<br/>any Overage for more than 14 days</p> | <p>Gravwell customer solutions manager will notify Customer via email of the Overage (the **“Second Notice”**) and provide the new annual license cost based on Customer’s Overages, prorated for the remainder of Customer’s then-current Subscription Term (the **“Overage Charges”**) should the Overage continue.</p> |
| Phase 3  | <p>***Following the Second Notice:***<br/>> 15% Overage within 45-day window</p> | <p>Customer’s designated Gravwell customer solutions manager will issue an invoice for the Overage Charges (the **“Overage Invoice”**).</p> | 
| Phase 4  | <p>***Following Issuance of the Overage Invoice:***<br/>Nonpayment 30 days from invoice date</p> | <p>Payment Failure Notifications shall be presented to users via the user-interface of the Software.</p> | 
| Phase 5  | <p>***Following Issuance of the Overage Invoice:***<br/>Nonpayment 45 days from invoice date.</p> | <p>Customer’s data ingestion shall be automatically limited to the Daily Ingestion Limits and excess data will be restricted from processing and/or may be deleted from the system, with no guarantees on which data will be dropped.</p> |

### 4.3. _Overage Disputes:_  
In the event of a good-faith dispute regarding Overage calculations or the amount due by Customer for any Overages, Customer must notify Gravwell in writing within 10 business days of receiving the applicable Overage report.  Both parties agree to work in good faith to resolve the dispute within 30 days. Customer will not be in breach of its payment obligations for Overages Charges (if any) under this Policy or the Agreement for disputes raised by Customer in good faith and in accordance with this Policy.

## 5. Temporary Capacity Increase:
### 5.1. _Request Process:_  
To initiate a request for a Temporary Capacity Increase, Customer must submit a written request to Gravwell at least 10 business days prior to the anticipated start of the one-off project. The request must include:
* The desired increase in daily ingestion capacity (in GB),
* The expected start and end dates of the project,
* A brief description of the project and its data ingestion requirements.

### 5.2. _Approval:_
Upon receiving Customer's request, Gravwell will evaluate the request and notify Customer if acceptable, following which, the parties shall execute an Order Form for the Temporary Capacity Increase that includes:
* The approved increase in the Daily Ingestion Limit (the **“Temporary Capacity Limits”**),
* The applicable fees for the Temporary Capacity Increase,
* The duration of the Temporary Capacity Increase (the **“Temporary Capacity Period”**),
* Any other relevant terms and conditions.

### 5.3. _Overages:_  
Overages of the Temporary Capacity Limits during the Temporary Capacity Period shall be subject to and addressed pursuant to Section 4.  For clarity, the term: (a) “Overages” as used herein shall include an Overage in the Temporary Capacity Limits during the Temporary Capacity Period, and (b) “Daily Ingestion Limits” as used in Section 4 shall refer to the Temporary Capacity Limits with respect to such Overages.

### 5.4. _Expiration:_  
Upon expiration of the Temporary Capacity Period, the Temporary Capacity Limits shall immediately terminate, and Customer's permitted volume limits shall automatically revert to the Daily Ingestion Limit.  

### 5.5. _Extension:_  
If Customer requires an extension of the Temporary Capacity Period, a written request must be submitted by Customer to Gravwell at least 5 business days prior to the end of the current Temporary Capacity Period. Gravwell will review the request and, if approved, extend the Temporary Capacity Period under the same terms or any revised terms as mutually agreed, and subject to Customer’s payment of the additional fees for such extended period.

### 5.6. _Additional Conditions:_  
Gravwell reserves the right to deny requests for a Temporary Capacity Increase if the request is deemed unreasonable or if accommodating the request would adversely affect Gravwell’s ability to deliver services to other customers.

## 6. Fees; Payment Terms:
All payments due hereunder, whether for Overage Charges or in connection with an Increase (including any Temporary Capacity Increase), shall be due and payable within 30 days of the applicable invoice date and made in accordance with the payment terms of the Agreement.

## 7. Notifications: 
Except as otherwise set forth herein, all notices to Customer under this Policy shall be sent to Customer’s designated contact specified in the Order Form via email, and shall be deemed delivered when sent (provided that Gravwell does not receive any error in transmission or failure to deliver notification).  In addition, Customer acknowledges and agrees that certain notifications may be sent via the user-interface of the Software as set forth in this Policy.  Customer consents to receiving electronic notifications from Gravwell via email and through the user-interface of the Software as described in this Policy.  Customer agrees that any notices, agreements, disclosures or other communications that Gravwell sends Customer electronically will satisfy any legal communication requirements, including that such communications be in writing, to the extent permitted by applicable law.

## 8. Modifications:
Gravwell may modify the terms of this Policy at any time; provided that, such modifications do not materially change the terms of this Policy during the Customer’s then-current Subscription Term. It is Customer’s responsibility to regularly review this Policy for updates. Customer may sign up to receive notices of updates to this Policy by emailing [legal@gravwell.io](mailto:legal@gravwell.io).  If Customer signs up for such updates, Gravwell will provide Customer reasonable prior notice (via email to Customer’s email address specified in the sign-up form) of any modifications to this Policy. Notwithstanding anything to the contrary herein, Gravwell will provide notice of any material changes to this Policy at least 90 days prior to the end of the then-current Subscription Term. For the avoidance of doubt, any such material changes shall not take effect until the start of Customer’s next Subscription Term. The most current version of this Policy is accessible at [https://docs.gravwell.io/sales/overages.html](https://docs.gravwell.io/sales/overages.html).
