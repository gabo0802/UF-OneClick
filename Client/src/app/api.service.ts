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

  login(userData: {password: string, username: string}): Observable<Object>{

    return this.http.post('/api/login', JSON.stringify(userData));
  }

  createUser(userData: {username: string, email: string, password: string}): Observable<Object>{   

    return this.http.post('/api/accountcreation', JSON.stringify(userData)); 
  }

  //UserInformation
  getUserInfo(): Observable<Object> {
    return this.http.get('api/alluserinfo');
  }

  //User subscriptions
  getActiveUserSubscriptions(): Observable<Subscription[]> {

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

  //All Subscriptions
  getAllSubscriptions(): Observable<Subscription[]> {

    return this.http.get('/api/subscriptions/services').pipe(
      map( (res: Object) => {

        let userSubs: Subscription[] = [];

        let data = JSON.stringify(res);
        let subData = JSON.parse(data);
        
        for(const sub of subData){        

          userSubs.push(sub);          
        }

        return userSubs;
      })
    );
  }

  getAllInactiveUserSubscriptions(): Observable<Subscription[]> {

    return this.http.get('/api/subscriptions').pipe(
      map( (res: Object) => {

        let userSubs: Subscription[] = [];

        let data = JSON.stringify(res);
        let subData = JSON.parse(data);
        
        for(const sub of subData){
          
          //Converts to a javascript date Object
          //Provisional as of now for easier displaying of date added in subscription table
          sub.dateadded = new Date(sub.dateadded);

          //checks to see if they have been removed, which means inactive
          if(sub.dateremoved !== ""){
            sub.dateremoved = new Date(sub.dateremoved);
            userSubs.push(sub); 
          }                   
        }
        
        return userSubs;
      })
    );
  }

  createUserSubscription(subName: string, subPrice: string): Observable<Object> {

    let subData = {name: subName, price: subPrice};

    return this.http.post('api/subscriptions/createsubscription', subData);
  }

  addUserSubscription(subName: string): Observable<Object> {

    let subData = {name: subName};

    return this.http.post('api/subscriptions/addsubscription', subData);
  }

  addOldUserSubscription(subData: {name: string, price: string, dateadded: string, dateremoved: string}): Observable<Object> {
    return this.http.post('api/subscriptions/addoldsubscription', subData);
  }

  deactivateSubscription(subName: string): Observable<Object> {

    let subData = {name: subName};

    return this.http.post('api/subscriptions/cancelsubscription', subData);
  }

  reactivateSubscription(subName: string): Observable<Object> {
    
    let subData = {name: subName};

    return this.http.post('/api/subscriptions/addsubscription', subData);
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