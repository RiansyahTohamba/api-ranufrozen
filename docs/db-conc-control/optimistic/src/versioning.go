package main
type TransactionRepository struct {
	db *sql.DB
}

func versioning() {
	// The VERSIONING OPTIMISTIC LOCKING option:
	- SELECT iD, val1, val2, version
		  FROM theTable
		  WHERE iD = @theId;
	- {code that calculates new values}
	- UPDATE theTable
		  SET val1 = @newVal1,
			  val2 = @newVal2,
			  version = version + 1
		  WHERE iD = @theId
			  AND version = @oldversion;
	- {if AffectedRows == 1 }
	-     {go on with your other code}
	- {else}
	-     {decide what to do since it has gone bad... in your code}
	- {endif}
	
}
