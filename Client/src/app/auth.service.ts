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

  userLogIn(): void {

    this.userLoggedIn = true;    
  }

  userLogOut(): void {

    this.userLoggedIn = false;
  }
}
