import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-add-subscription',
  templateUrl: './add-subscription.component.html',
  styleUrls: ['./add-subscription.component.css']
})
export class AddSubscriptionComponent {

  addSubsriptionForm = new FormGroup({
    'name': new FormControl(null, [Validators.required, Validators.pattern('^[ A-z0-9]+$')]),
    'price': new FormControl(null, Validators.required)
  });

}
