import { Component } from '@angular/core';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent {
  todayDate: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());
  displayHome: boolean = true;
  displayReset: boolean = false;
  displayNewsletter: boolean = false;

  goHome(): void {

    if(this.displayHome !== true){

      this.displayReset = false;
      this.displayHome = true;
      this.displayNewsletter = false;

    }
  }

  showReset(): void {
    if(this.displayReset !== true){

      this.displayReset = true;
      this.displayHome = false;
      this.displayNewsletter = false;
    }    
  }

  showNewsletter(): void {
    if(this.displayNewsletter !== true){

      this.displayNewsletter = true;
      this.displayHome = false;
      this.displayReset = false;
    }
  }
}
