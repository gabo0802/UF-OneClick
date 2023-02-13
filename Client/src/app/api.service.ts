import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  public output: string = '';

  constructor(private http: HttpClient) { }

  
  Login(userData: {password: string, username: string}){
    //document.cookie = "loginOutput=none; path=/; max-age=100" 

    this.http.post('/api/login', JSON.stringify(userData)).subscribe((res)=> {
      //console.log("Inside .subscribe")
      this.output = document.cookie
      console.log(res)
      console.log(document.cookie)

      //console.log(userData.username);
      //console.log(userData.password);
      //console.log(userData.email);
    });

    //console.log("Outside .subscribe") //different value
    //console.log(document.cookie)

    /*this.http.post('http://localhost:8000/', JSON.stringify(userData)).subscribe((res)=>{
      console.log(res);
    });*/
  }

  createUser(userData: {email: string, password: string, username: string}): Observable<Object>{
    document.cookie = "signupOutput=none; path=/; max-age=100" 

    return this.http.post('/api/accountcreation', JSON.stringify(userData));

  }

  public getOutput() {
    if (this.output == '' || this.output == 'signupOutput=none'){
      this.output = 'Please enter the required information below!'
    }
    return this.output;
  }
}
