import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  public output: string = '';

  constructor(private http: HttpClient) { }

  

  createUser(userData: {email: string, password: string, username: string}){
    document.cookie = "signupOutput=none; path=/; max-age=100" 

    this.http.post('/api/accountcreation', JSON.stringify(userData)).subscribe((res)=> {
      this.output = document.cookie
      console.log(res)
      //console.log(document.cookie)
      console.log(this.output)

      //console.log(userData.username);
      //console.log(userData.password);
      //console.log(userData.email);
    });

    /*this.http.post('http://localhost:8000/', JSON.stringify(userData)).subscribe((res)=>{
      console.log(res);
    });*/

  }

  public getOutput() {
    return this.output;
  }
}
