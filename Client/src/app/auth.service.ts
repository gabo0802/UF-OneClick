import { Injectable } from '@angular/core';


@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private userLoggedIn: boolean = false;

  constructor() { }

  isLoggedIn(): boolean {

    if(document.cookie.includes("currentUserID=")){
      this.userLogIn();
    }

    return this.userLoggedIn;
  }

  isAdmin(): boolean {

    if(document.cookie.includes("currentUserID=1")){
      return true;
    }
    else {
      return false;
    }
  }

  userLogIn(): void {

    this.userLoggedIn = true;    
  }

  userLogOut(): void {

    if(this.userLoggedIn && document.cookie.includes("currentUserID=")){
      
      document.cookie = "currentUserID=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/";
    }

    this.userLoggedIn = false;
  }
}
