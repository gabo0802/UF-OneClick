import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit, ViewChild } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';
import { MatAccordion } from '@angular/material/expansion';
import { forkJoin } from 'rxjs';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit{

  @ViewChild(MatAccordion) accordion: MatAccordion;

  constructor(private api: ApiService, private dialogs: DialogsService) {
    this.accordion = new MatAccordion()
    this.accordion.openAll();
  }

  username: string = '';
  subscriptionList: Subscription[] = [];

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
}