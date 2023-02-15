import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  constructor(private router: Router) {};  

  ngOnInit(){
    if (document.cookie.includes("currentUserID=")){
        this.router.navigate(['users']);
    }
  }

  onSubmit(){
    console.log("Test 2")
  }
}
