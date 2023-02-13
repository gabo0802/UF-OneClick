import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  loginForm: FormGroup = {} as FormGroup;
  constructor(private api: ApiService) {};
  hide: boolean = true;

  ngOnInit(){
    this.loginForm = new FormGroup({
      'username': new FormControl(null, null),
      'password': new FormControl(null, null)
    });
  }

  onSubmit(){
    this.api.Login(this.loginForm.value).subscribe( (res: Object) => {
  
      const response: string = JSON.stringify(res);
      const responseMessage = JSON.parse(response);

      //Incorrect Password or Database Error
      if(responseMessage["Error"] !== undefined){        
        console.log(responseMessage["Error"])
      }//Logged in and User Subscriptions returned
      else{
        //might be able to do with for loop but so far just got this to work with while only
        let index: number = 0;
        let test: string = "";
        var allSubscriptions = [];

        while (responseMessage[index] != null){
          console.log(responseMessage[index]); //get Subcription
          console.log(responseMessage[index]["name"]); //get value from Subcription (see userData for all parameters)

          test = responseMessage[index]["name"];
        
          allSubscriptions.push(responseMessage[index]);

          index += 1;
        }

        console.log("Subscription Name: " + test); //testing string variable
        console.log(allSubscriptions); //testing array
        console.log(document.cookie); //showing created cookie (possibly for future use)
      }
    });

    this.loginForm.reset();
  }

}
