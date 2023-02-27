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

  // post_request(userData: {name: string}, url:string): Observable<Object>{
  //       return this.http.post(url, JSON.stringify(userData));
  // }

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

  getEmailandUsername(): Observable<Array<string>> {
    return this.http.get<{[key: string]: string, message: string}>('api/userinfo').pipe(
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

  updateUsername(newUsername: string): Observable<Object> {
    const newUserUsername = {username: newUsername};
    return this.http.put('api/changeusername', newUserUsername);
  }

  updateUserEmail(userEmail: string): Observable<Object> {
    const newUserEmail = {email: userEmail};
    return this.http.put('api/changeemail', newUserEmail);
  }

  updateUserPassword(userOldPassword: string, userNewPassword: string): Observable<Object> {
    const passwords= {oldPassword: userOldPassword, newPassword: userNewPassword};
    return this.http.put('api/changepassword', passwords);
  }

  deleteUserAccount(): Observable<Object> {
    return this.http.delete('api/deleteuser');
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

}
