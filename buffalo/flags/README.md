## Quirks

- Buffalo was missing a package from the start which caused the initial start to fail. Ref: https://github.com/gobuffalo/buffalo/issues/1622

### Create Your Databases

Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

	$ buffalo pop create

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.
