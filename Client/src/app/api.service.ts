import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) { }


  createUser(userData: {email: string, password: string, username: string}){
    
    this.http.post('/api/accountcreation', JSON.stringify(userData)).subscribe((res)=>{
      console.log(res);

      console.log("Hello World");

      console.log(userData.username);
      console.log(userData.password);
      console.log(userData.email);

    }); //Can't get parameters

    /*this.http.post('http://localhost:8000/', JSON.stringify(userData)).subscribe((res)=>{
      console.log(res);
    });*/

  }
}
