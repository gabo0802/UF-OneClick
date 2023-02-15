import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  constructor(private http: HttpClient, private router: Router) {};  

  ngOnInit(){
    if (document.cookie.includes("currentUserID=")){
        console.log(document.cookie)
        this.router.navigate(['users']);
    }
  }

}
