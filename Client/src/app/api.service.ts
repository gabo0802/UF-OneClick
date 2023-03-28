import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs';
import { Subscription } from './subscription.model';

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

  //UserInformation
  getUserInfo(): Observable<Object> {
    return this.http.get('api/alluserinfo');
  }

  //User subscriptions
  getUserSubscriptions(): Observable<Subscription[]> {

    return this.http.get('/api/subscriptions/active').pipe(
      map( (res: Object) => {

        let userSubs: Subscription[] = [];

        let data = JSON.stringify(res);
        let subData = JSON.parse(data);
        
        for(const sub of subData){
          
          //Converts to a javascript date Object
          //Provisional as of now for easier displaying of date added in subscription table
          sub.dateadded = new Date(sub.dateadded);

          userSubs.push(sub);          
        }
        
        return userSubs;
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

  updateTimezone(newTimezone: string): Observable<Object> {
    const timezone = {timezonedifference: newTimezone};
    return this.http.put('api/changetimezone', timezone);
  }

  deleteUserAccount(): Observable<Object> {
    return this.http.delete('api/deleteuser');
  }
}