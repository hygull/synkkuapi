views.ShowSuccessOrErrorAsJSON(ctx, "InvalidAppSignKey", "App signin key is invalid", 101, 1)

views.ShowSuccessOrErrorAsJSON(ctx, "InvalidEmailFormatError", 
	"The enetered email is not valid", 102, 2)
	
views.ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", 
	"email & app_sign_in_key keys, both are mandatory for existing users, 
	check if you have provided  wrong key names", 103, 3)

views.ShowSuccessOrErrorAsJSON(ctx, "NoPostDataError", 
	"You haven't sent the POST data", 104, 4)

views.ShowSuccessOrErrorAsJSON(ctx, "PostMethodNotFoundError", 
	"The data is not being sent using POST method", 107, 7)


views.ShowSuccessOrErrorAsJSON(rw, "InsertQueryExecutionError", 
	"Error in executing the insert query", 108, 8)

views.ShowSuccessOrErrorAsJSON(rw, "SelectQueryExecutionError", 
	"Error in executing the select query", 109, 9)

views.ShowSuccessOrErrorAsJSON(rw, "UpdateQueryExecutionError", 
	"Error in execution of update query", 110, 10)

views.ShowSuccessOrErrorAsJSON(rw, "DeleteQueryExecutionError", 
	"Error in executing the delete query", 111,11)

views.ShowSuccessOrErrorAsJSON(ctx, "DBConnectionTestError", 
	"Error in connection with database", 112, 12)

views.ShowSuccessOrErrorAsJSON(ctx, "AccountRecreationAttemptError", 
	"This email ID is already registered...so please login", 113, 13)

views.ShowSuccessOrErrorAsJSON(ctx, "InvalidFirstLoginAppSignKeyError", 
	"You have not specified the proper App SignIn Key required for first login. 
	Demand it from your App sevice provider", 114, 14)

