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

  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit(): void {
    
    if(this.authService.isLoggedIn()){
      this.router.navigate(['users']);
    }
  }
}
