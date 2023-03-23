import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Subscription } from 'src/app/subscription.model';

@Component({
  selector: 'app-edit-subscription',
  templateUrl: './edit-subscription.component.html',
  styleUrls: ['./edit-subscription.component.css']
})
export class EditSubscriptionComponent implements OnInit{

  constructor(@Inject(MAT_DIALOG_DATA) public data: {subData: Subscription}) { }

  ngOnInit(): void {
    console.log();
  }

}
