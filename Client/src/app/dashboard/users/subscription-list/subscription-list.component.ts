import { Component, Input } from '@angular/core';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';

@Component({
  selector: 'app-subscription-list',
  templateUrl: './subscription-list.component.html',
  styleUrls: ['./subscription-list.component.css']
})
export class SubscriptionListComponent {

  constructor(private dialogs: DialogsService) {}

  @Input() subscriptionList: Subscription[] = [];

  displayedColumns: string[] = ['sub-name', 'sub-price', 'sub-actions'];

  addSubscription(): void {

    this.dialogs.addSubscription();
  }

  editSubscription(subscriptionInfo: Subscription): void {

    this.dialogs.editSubscription(subscriptionInfo);
  }
}