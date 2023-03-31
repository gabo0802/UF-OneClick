import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit, ViewChild } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';
import { MatAccordion } from '@angular/material/expansion';

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

    this.api.getUserInfo().subscribe({

      next: (res: Object) => {

        let data = JSON.stringify(res);
        let userData = JSON.parse(data);

        this.username = userData.username;
      },
      error: (error: HttpErrorResponse) => {

        this.dialogs.errorDialog("Unexpected Error getting user data!", "Please try again later.");
      }
    });

    this.api.getActiveUserSubscriptions().subscribe({
      
      next: (res: Subscription[]) => {       

        this.subscriptionList = res;
        
      },
      error: (error: HttpErrorResponse) => {

        this.dialogs.errorDialog("Unexpected Error!", error["error"]["Error"] + " Please try again later.");
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