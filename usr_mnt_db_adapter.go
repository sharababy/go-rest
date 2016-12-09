//user management database adapters


package main


import(
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"log"
)


var err_insert chan error

func Insert_User(DB_Name string , Collection_Name string, session *mgo.Session, data User ) <-chan error{

go func() {

	c :=session.DB(DB_Name).C(Collection_Name)

	err := c.Insert(data)

	err_insert <- err

	if(err!=nil){
		log.Println(err)
	}
	//fmt.Println("Inserted doc..")
	defer session.Close()


	}()
	return err_insert
}


func DumpAll( DB_Name string, Collection_Name string, session *mgo.Session ) ([]User , error) {

	var allUsers []User

	f := session.DB(DB_Name).C(Collection_Name)


	err := f.Find(nil).All(&allUsers)


		if(err!=nil){
        	
        		return allUsers , err 
        	} else{
        		return allUsers , nil
        

    		}

}


func Find_User( DB_Name string, Collection_Name string, session *mgo.Session,find_type string ,find_with string) (User,error) {

		var lookfor User

			f := session.DB(DB_Name).C(Collection_Name)

			FindWith := bson.M{find_type : find_with}

        	err := f.Find(FindWith).One(&lookfor)

        	if(err!=nil){
        	
        		return lookfor , err 
        	} else{
        		return lookfor , nil
        

    		}
}

func Delete_User( DB_Name string, Collection_Name string, session *mgo.Session,find_type string ,find_with string) (error) {


       	collection := session.DB(DB_Name).C(Collection_Name)


		Delete_this := bson.M{ find_type : find_with }

        err := collection.Remove(Delete_this)

        if(err!=nil){
        	return err 
        } else{
        	return nil
        }
}


func Update_User(DB_Name string, Collection_Name string, session *mgo.Session,find_type string ,find_with string ,new_data User) (User,error){
	
	UpdateWith := bson.M{ find_type : find_with }
	
	old_data,err := Find_User(DB_Name,
						Collection_Name,
						session,
						find_type,
						find_with)

	fmt.Println(old_data)

	if(err!=nil){
		return new_data,err
	
	} else{
		// this is done to ensure that the if any field is left blank in JSON
		// it should not be left blank in the db, it should be replaced by the 
		//old values in the db
		if(new_data.UserId==""){ 
				new_data.UserId = old_data.UserId;
		}

		if(new_data.FullName==""){
				new_data.FullName = old_data.FullName;
		}

		if(new_data.Phone==""){
				new_data.Phone = old_data.Phone;
		}
		if(new_data.Email==""){
				new_data.Email = old_data.Email;
		}
		if(new_data.ClientName==""){
				new_data.ClientName = old_data.ClientName;
		}
		if(new_data.Status==""){
				new_data.Status = old_data.Status;
		}
	}

	change := bson.M{"$set": bson.M{
		"UserId" : new_data.UserId,
		"Phone" : new_data.Phone,
		"Email" : new_data.Email,
		"FullName" : new_data.FullName,
		"Status" : new_data.Status,
		"ClientName" : new_data.ClientName,
		},	 
	}

	c :=session.DB(DB_Name).C(Collection_Name)

	err = c.Update(UpdateWith, change)

	if(err!=nil){
		log.Println(err)
		return new_data,err
	}

	return new_data, nil
}



func CreateSession(DB_Url string) *mgo.Session{

	session, err := mgo.Dial(DB_Url)

	if(err!=nil){
		log.Fatal(err)
	}

	session.SetSafe(&mgo.Safe{})

	return session
}