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

type User struct{ 

	Name string `bson:"Name" json:"Name" `
	Phone string `bson:"Phone" json:"Phone" `
	Email string `bson:"Email" json:"Email" `

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

    session := CreateSession("mongodb://localhost:27017") 

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

    session := CreateSession("mongodb://localhost:27017")


    //Update_User is a CRUD helper function that is defined in the db_adapter.go file
    //it take db name , collection name , session and find_with and find_typr , User struct as input
    // it returns type User and type error

   	new_user , errr := Update_User("wandr","endusers",session,find_type, find_with,new_user)

   	if(errr!=nil){
   		fmt.Fprintf(w,"Error, Unable to update field.")
   		fmt.Println(errr)
   	} else{
    	fmt.Fprintf(w,"Successfully Updated")

    	fmt.Println("New details - \n Name - ",new_user.Name,"\n Phone - ", new_user.Phone,"\n Email - ",new_user.Email)
	}
}


func DeleteJSON(w http.ResponseWriter , r *http.Request , p httprouter.Params){

	find_type := p.ByName("find_type") // it extracts the value of find_type paramater from the url

	find_with := p.ByName("find_with") // it extracts the value of find_with paramater from the url

    //CreateSession() is defined in the db_adapter.go file
    //it returns a *mgo.Session type 
    //it establishes a session between the MongoDB instance and the golang app

    session := CreateSession("mongodb://localhost:27017")


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


func FindJSON(w http.ResponseWriter , r *http.Request , p httprouter.Params){

	find_type := p.ByName("find_type") // it extracts the value of find_type paramater from the url

	find_with := p.ByName("find_with") // it extracts the value of find_with paramater from the url

	decoder := json.NewDecoder(r.Body) //  JSON decoder  for request.Body to convert to JSON

	var Find_user User

    err := decoder.Decode(&Find_user)

    if err != nil {	// error check for Decoder
        panic(err)
    }

    defer r.Body.Close()

    //CreateSession() is defined in the db_adapter.go file
    //it returns a *mgo.Session type 
    //it establishes a session between the MongoDB instance and the golang app

    session := CreateSession("mongodb://localhost:27017")


    //Find_User is a CRUD helper function that is defined in the db_adapter.go file
    //it take db name , collection name , session and find_with and find_typr , User struct as input
    // it returns type User and type error

   	Find_user , errr := Find_User("wandr","endusers",session,find_type, find_with)

   	if(errr!=nil){
   		fmt.Fprintf(w,"Error, Unable to find user.")
   		fmt.Println(errr)
   	} else{
    	fmt.Fprintf(w,"Successfully found user:\n");
    	fmt.Fprintf(w,"User details - \n Name - %s \n Phone - %s \n Email - %s",Find_user.Name,Find_user.Phone,Find_user.Email)

    	fmt.Println("User details - \n Name - ",Find_user.Name,"\n Phone - ", Find_user.Phone,"\n Email - ",Find_user.Email)
	}
}


/*
func Home(w http.ResponseWriter , r *http.Request){

			ip := r.RemoteAddr
			file := r.URL.Path
			filename:=file[1:]



			if len(filename)==0 {
				ServeHome(w,r)
				fmt.Println(ip," requested home page...")


			/* }  else if filename == "register" {
				message := ""
				Register(w,r,nil,message)

			
			} else{  
				http.ServeFile(w,r,filename+".html")
				fmt.Println(ip," requested",filename,"page...")	
			}
	

}




func ServeHome(w http.ResponseWriter, r *http.Request){

		

					t,_ := template.ParseFiles("usr_mnt_addusr.html")
		

					t.Execute(w,nil)

	}

	*/