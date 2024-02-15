package html

func WelcomePageHTML() string {
    return `
		<!DOCTYPE html>
		<html lang="en">
			<head>
			    <meta charset="UTF-8">
			    <meta name="viewport" content="width=device-width, initial-scale=1.0">
			    <title>Profitability</title>
			</head>
		    <body>
		        <h1>Welcome to the Profitability API</h1>
		        <p>Thank you for visiting Profitability! For more information, see the <a href="https://rmottanet.gitbook.io/profitability">documentation</a>.</p>
		    </body>
		</html>
    `
}
