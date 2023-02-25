import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {  

  constructor(private http: HttpClient) { }

  post_request__with_data(userData: {username: string, email: string, password: string, name: string, price: string, dateadded: string, dateremoved: ""}, url:string): Observable<Array<string>>{
     return this.http.post<{[key: string]: string, message: string}>(url, JSON.stringify(userData)).pipe(
        map( (statusMessage) => {        

          const resultMessage: string[] = [];
          
          for(const key in statusMessage){

            resultMessage.push(key);
            resultMessage.push(statusMessage[key]);

          }       

          return resultMessage;
        })
      );
  }

  post_request(userData: {name: string}, url:string): Observable<Object>{
        return this.http.post(url, JSON.stringify(userData));
  }

  /*login(userData: {password: string, username: string}): Observable<Array<string>>{

    return this.http.post<{[key: string]: string, message: string}>('/api/login', JSON.stringify(userData)).pipe(
      map( (statusMessage) => {        

        const resultMessage: string[] = [];
        
        for(const key in statusMessage){

          resultMessage.push(key);
          resultMessage.push(statusMessage[key]);

        }       

        return resultMessage;
      })
    );
  }

  createUser(userData: {username: string, email: string, password: string}): Observable<Array<string>>{   

    return this.http.post<{[key: string]: string, message: string}>('/api/accountcreation', JSON.stringify(userData)).pipe(
      map( (statusMessage) => {        

        const resultMessage: string[] = [];
        
        for(const key in statusMessage){

          resultMessage.push(key);
          resultMessage.push(statusMessage[key]);

        }       

        return resultMessage;
      })
    );
  }*/

  /*public getOutput() {
    if (this.output == '' || this.output == 'signupOutput=none'){
      this.output = 'Please enter the required information below!'
    }
    return this.output;
  }*/
}
