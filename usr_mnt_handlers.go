//This file belongs to a user management system for all kinds for users.

// this file handles the http function handlers

package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"httprouter"
	//"html/template"

)

//	Every User can be fit into this common data type knows as User
// More meta data to be added here ::..>

const db string = "mongodb://user1:12345678@ds119738.mlab.com:19738/wandr" //"mongodb://localhost:27017"



type User struct{ 

  UserId string `bson:"UserId" json:"UserId" `
	FullName string `bson:"FullName" json:"FullName" `
	Phone string `bson:"Phone" json:"Phone" `
	Email string `bson:"Email" json:"Email" `
  ClientName string `bson:"ClientName" json:"ClientName" `
  Status string `bson:"Status" json:"Status" `
  
}

//This function receives the data in application/json encoding 
//and then feeds it into the database by calling the Insert_User() function

//it also prints if the transation was successful or not


func ReceiveJSON(w http.ResponseWriter , r *http.Request , _ httprouter.Params){

	decoder := json.NewDecoder(r.Body) // this is a JSON decoder

    var new_user User  // a new User data type for new user data

    err := decoder.Decode(&new_user) 


    if err != nil { // error handling for JSON decoder
        panic(err)
    }

    defer r.Body.Close() // closing of the request.body tag

    //CreateSession() is defined in the db_adapter.go file
    //it returns a *mgo.Session type 
    //it establishes a session between the MongoDB instance and the golang app

    session := CreateSession(db) 

    //Insert_User is a CRUD helper function that is defined in the db_adapter.go file
    //it take db name , collection name , session and User struct as input
    // it returns type error
    err = Insert_User("wandr","endusers",session,new_user)

    if(err!=nil){
   		fmt.Fprintf(w,"Error, Unable to insert field")
   	
   	}else{
    	//fmt.Println("New details - \n Name - ",new_user.Name,"\n Phone - ", new_user.Phone,"\n Email - ",new_user.Email)
	}
}


//This function upadates the data in application/json encoding 
//and then feeds it into the database by calling the Update_User() function
//It is the basis of the REST API .

// It is the funtion call for 
// URL : "/update/:find_type/:find_with"

//it also printd if the transation was successful or not

func UpdateJSON(w http.ResponseWriter , r *http.Request , p httprouter.Params){

	find_type := p.ByName("find_type") // it extracts the value of find_type paramater from the url

	find_with := p.ByName("find_with") // it extracts the value of find_with paramater from the url

	decoder := json.NewDecoder(r.Body) //  JSON decoder  for request.Body to convert to JSON

    var new_user User 

    err := decoder.Decode(&new_user)

    if err != nil {	// error check for Decoder
        panic(err)
    }

    defer r.Body.Close()

    //CreateSession() is defined in the db_adapter.go file
    //it returns a *mgo.Session type 
    //it establishes a session between the MongoDB instance and the golang app

    session := CreateSession(db)


    //Update_User is a CRUD helper function that is defined in the db_adapter.go file
    //it take db name , collection name , session and find_with and find_typr , User struct as input
    // it returns type User and type error

   	new_user , errr := Update_User("wandr","endusers",session,find_type, find_with,new_user)

   	if(errr!=nil){
   		fmt.Fprintf(w,"Error, Unable to update field.")
   		fmt.Println(errr)
   	} else{
    	fmt.Fprintf(w,"Successfully Updated")

    	fmt.Println("New details - \n FullName - ",new_user.FullName,"\n Phone - ", new_user.Phone,"\n Email - ",new_user.Email,"\n UserId - ",new_user.Email,"\n ClientName - ",new_user.ClientName,"\n Status - ",new_user.Status) 
	}
}


func DeleteJSON(w http.ResponseWriter , r *http.Request , p httprouter.Params){

	find_type := p.ByName("find_type") // it extracts the value of find_type paramater from the url

	find_with := p.ByName("find_with") // it extracts the value of find_with paramater from the url

    //CreateSession() is defined in the db_adapter.go file
    //it returns a *mgo.Session type 
    //it establishes a session between the MongoDB instance and the golang app

    session := CreateSession(db)


    //Delete_User is a CRUD helper function that is defined in the db_adapter.go file
    //it take db name , collection name , session and find_with and find_type as input
    // it returns type error

   	err := Delete_User("wandr","endusers",session,find_type, find_with)

   	if(err!=nil){
   		fmt.Fprintf(w,"Error, Unable to delete field.")
   		fmt.Println(err)
   	} else{
    	fmt.Fprintf(w,"Successfully Deleted")
	}
}

func DumpJSON(w http.ResponseWriter , r *http.Request , p httprouter.Params){


        session := CreateSession(db)


           var all []User

        all , err := DumpAll("wandr","endusers",session)

         if(err!=nil){
                
                fmt.Fprintf(w,"Error, Unable to find user.")
                fmt.Println(err)
        
          } else{
      
                w.Header().Set("Content-Type", "application/json")                 

                user, errr := json.Marshal(all)
                if(errr!=nil){
                  fmt.Println(errr);
                }
                w.Write(user)

                fmt.Println(user)
        }
}


func FindJSON(w http.ResponseWriter , r *http.Request , p httprouter.Params){

	find_type := p.ByName("find_type") // it extracts the value of find_type paramater from the url

	find_with := p.ByName("find_with") // it extracts the value of find_with paramater from the url

    //CreateSession() is defined in the db_adapter.go file
    //it returns a *mgo.Session type 
    //it establishes a session between the MongoDB instance and the golang app

    session := CreateSession(db)

    
    

    //Find_User is a CRUD helper function that is defined in the db_adapter.go file
    //it take db name , collection name , session and find_with and find_typr , User struct as input
    // it returns type User and type error

   	Find_user , errr := Find_User("wandr","endusers",session,find_type, find_with)

   	if(errr!=nil){
   		fmt.Fprintf(w,"Error, Unable to find user.")
   		fmt.Println(errr)
   	} else{

      w.Header().Set("Content-Type", "application/json")  
    	
      user, err := json.Marshal(Find_user)

      if(err!=nil){
        fmt.Println(err)
      }
      w.Write(user)

    	fmt.Println(user)
	}



}
