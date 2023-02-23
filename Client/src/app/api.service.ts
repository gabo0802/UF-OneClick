import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {  

  constructor(private http: HttpClient) { }

  login(userData: {password: string, username: string}): Observable<Array<string>>{

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
  }


  // addUserSub(userData: {name: string}): Observable<Array<string>>{   

  //   return this.http.post<{[key: string]: string, message: string}>('/api/subscriptions/addsubscription', JSON.stringify(userData)).pipe(
  //     map( (statusMessage) => {        

  //       const resultMessage: string[] = [];
        
  //       for(const key in statusMessage){

  //         resultMessage.push(key);
  //         resultMessage.push(statusMessage[key]);

  //       }       

  //       return resultMessage;
  //     })
  //   );
  // }

  // removeUserSub(userData: {name: string}): Observable<Array<string>>{   

  //   return this.http.post<{[key: string]: string, message: string}>('/api/subscriptions/cancelsubscription', JSON.stringify(userData)).pipe(
  //     map( (statusMessage) => {        

  //       const resultMessage: string[] = [];
        
  //       for(const key in statusMessage){

  //         resultMessage.push(key);
  //         resultMessage.push(statusMessage[key]);

  //       }       

  //       return resultMessage;
  //     })
  //   );
  // }

  // addOldUserSub(userData: {name: string}): Observable<Array<string>>{   

  //   return this.http.post<{[key: string]: string, message: string}>('/api/subscriptions/addoldsubscription', JSON.stringify(userData)).pipe(
  //     map( (statusMessage) => {        

  //       const resultMessage: string[] = [];
        
  //       for(const key in statusMessage){

  //         resultMessage.push(key);
  //         resultMessage.push(statusMessage[key]);

  //       }       

  //       return resultMessage;
  //     })
  //   );
  // }

  // createSub(userData: {name: string, price: string}): Observable<Array<string>>{   
  //   userData.price = userData.price.replaceAll("$", "")
  //   console.log("test")

  //   return this.http.post<{[key: string]: string, message: string}>('/api/subscriptions/createsubscription', JSON.stringify(userData)).pipe(
  //     map( (statusMessage) => {        

  //       const resultMessage: string[] = [];
        
  //       for(const key in statusMessage){

  //         resultMessage.push(key);
  //         resultMessage.push(statusMessage[key]);

  //       }       

  //       return resultMessage;
  //     })
  //   );
  // }

  // getSubs(): Observable<Object>{
  //   return this.http.post('/api/subscriptions', null);
  // }

  // logout(): Observable<Object>{
  //   return this.http.post('/api/logout', null);
  // }

  // resetall(): Observable<Object>{
  //   return this.http.post('/api/reset', null);
  // }

  // sendNews(userData: {name: string}): Observable<Object>{
  //   return this.http.post('/api/news', JSON.stringify(userData));
  // }

  // getAllUserData(): Observable<Array<string>>{   
  //   return this.http.post<{[key: string]: string, message: string}>('/api/alldata', null).pipe(
  //     map( (statusMessage) => {        

  //       const resultMessage: string[] = [];
        
  //       for(const key in statusMessage){

  //         resultMessage.push(key);
  //         resultMessage.push(statusMessage[key]);

  //       }       

  //       return resultMessage;
  //     })
  //   );
  // }

  /*public getOutput() {
    if (this.output == '' || this.output == 'signupOutput=none'){
      this.output = 'Please enter the required information below!'
    }
    return this.output;
  }*/
}
