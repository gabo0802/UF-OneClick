import { HttpErrorResponse } from '@angular/common/http';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { result } from 'cypress/types/lodash';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';

@Component({
  selector: 'app-subscription-list',
  templateUrl: './subscription-list.component.html',
  styleUrls: ['./subscription-list.component.css']
})
export class SubscriptionListComponent {

  constructor(private dialogs: DialogsService, private api: ApiService) {}

  @Input() subscriptionList: Subscription[] = [];
  @Output() updateSubscriptions = new EventEmitter<boolean>();

  displayedColumns: string[] = ['sub-name', 'sub-price', 'sub-dateAdded', 'sub-actions'];
  
  addSubscription(): void {
    this.dialogs.addSubscription().afterClosed().subscribe((res: {isCreated: boolean, name: string}) => {


      if(res.isCreated === true){

        this.api.addUserSubscription(res.name).subscribe({

          next: (res: Object) => {
            this.updateSubscriptions.emit(true);
          },
          error: (error: HttpErrorResponse) => {
            this.dialogs.errorDialog("Error Adding Subscription!", "There was an erroring adding your subscription. Please try again later.")
          }
        });

      }
    });
  }

}