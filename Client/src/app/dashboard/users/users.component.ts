import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';
import { forkJoin } from 'rxjs';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit{

  constructor(private api: ApiService, private dialogs: DialogsService) { }

  username: string = '';
  subscriptionList: Subscription[] = [];
  todayDate: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());
  displaySubscriptions: boolean = true;
  displayReport: boolean = false;
  displayGraph: boolean = false;

  ngOnInit(): void {

    forkJoin({
      userInfo: this.api.getUserInfo(),
      activeSubs: this.api.getActiveUserSubscriptions(),
    }).subscribe({

      next: (res) => {
        this.subscriptionList = res.activeSubs;
        
        let data = JSON.stringify(res.userInfo);
        let userData = JSON.parse(data);
        this.username = userData.username;
      },
      error: (error: HttpErrorResponse) => {
        this.dialogs.errorDialog("Unexpected Error!", "Please try again later " + error["error"]["Error"]);
      }
    });    
  }

  updateActiveSubscriptions(update: boolean): void {

    if(update){

      this.api.getActiveUserSubscriptions().subscribe({
      
        next: (res: Subscription[]) => {       
  
          this.subscriptionList = res;
          
        },
        error: (error: HttpErrorResponse) => {
  
          this.dialogs.errorDialog("Unexpected Error!", error["error"]["Error"] + " Please try again later.");
        }
      });
    }
  }

  updateInactiveSubscriptions(update: boolean): void {

    if(update){

      this.api.getAllInactiveUserSubscriptions().subscribe({
      
        next: (res: Subscription[]) => {       
  
          this.subscriptionList = res;
          
        },
        error: (error: HttpErrorResponse) => {
  
          this.dialogs.errorDialog("Unexpected Error!", error["error"]["Error"] + " Please try again later.");
        }
      });
    }
  }

  showReport(): void {

    if(this.displayReport !== true){

      this.displaySubscriptions = false;
      this.displayReport = true;
      this.displayGraph = false;
    }
  }

  showSubscriptions(): void {

    //makes sure subscriptions aren't already displaying
    if(this.displaySubscriptions !== true){

      this.displaySubscriptions = true;
      this.displayReport = false;
      this.displayGraph = false;

      this.updateActiveSubscriptions(true);
    }    
  }

  showGraph(): void {

    //makes sure graph isn't already displaying
    if(this.displayGraph !== true){

      this.displayGraph = true;
      this.displaySubscriptions = false;
      this.displayReport = false;
    }
  }
}