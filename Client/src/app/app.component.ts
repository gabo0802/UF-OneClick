import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from './auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['app.component.css']
})
export class AppComponent {

  title = 'UF-OneClick';
  myHeaderState: HeaderState = 1;

  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit(): void {
    
    if(this.authService.isLoggedIn() && this.authService.isAdmin() == false){
      this.router.navigate(['users']);
    }
    else if(this.authService.isLoggedIn() && this.authService.isAdmin() == true){
      this.router.navigate(['admin']);
    }
  }

  setHeader() {
    let path = this.router.url.split('/')[1];

    switch(path) {
      case "":
        this.myHeaderState = 1;
        break;
      case "signup":
        this.myHeaderState = 2;
      break;
      case "login":
        this.myHeaderState = 2;
      break;
      case "users":
        this.myHeaderState = 3;
      break;
      case "admin":
        this.myHeaderState = 3;
        break;
      default:
        this.myHeaderState = 1;
        console.log("An error occurred and you were sent to the main page!");
    }
    
  }

}

enum HeaderState {
  LandingPage = 1,
  SignOrLogInPage = 2,
  UserPage = 3
}
