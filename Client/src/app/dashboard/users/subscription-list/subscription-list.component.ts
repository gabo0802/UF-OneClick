import { Component, Input } from '@angular/core';
import { Subscription } from 'src/app/subscription.model';

@Component({
  selector: 'app-subscription-list',
  templateUrl: './subscription-list.component.html',
  styleUrls: ['./subscription-list.component.css']
})
export class SubscriptionListComponent {

  @Input() subscriptionList: Subscription[] = [];

  displayedColumns: string[] = ['sub-name', 'sub-price', 'sub-edit'];
}
