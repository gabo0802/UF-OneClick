import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  //public output: string = '';

  constructor(private http: HttpClient) { }

  Login(userData: {password: string, username: string}): Observable<Object>{    
    return this.http.post('/api/login', JSON.stringify(userData));
  }

  createUser(userData: {email: string, password: string, username: string}): Observable<Object>{    

    return this.http.post('/api/accountcreation', JSON.stringify(userData));
  }

  /*public getOutput() {
    if (this.output == '' || this.output == 'signupOutput=none'){
      this.output = 'Please enter the required information below!'
    }
    return this.output;
  }*/
}
