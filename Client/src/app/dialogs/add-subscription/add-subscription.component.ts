import { HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';

@Component({
  selector: 'app-add-subscription',
  templateUrl: './add-subscription.component.html',
  styleUrls: ['./add-subscription.component.css']
})
export class AddSubscriptionComponent {

  constructor(private api: ApiService, private dialogs: DialogsService, private dialogRef: MatDialogRef<AddSubscriptionComponent>) {}

  addSubsriptionForm = new FormGroup({
    'name': new FormControl(null, [Validators.required, Validators.pattern('^[ A-z0-9]+$')]),
    'price': new FormControl(null, Validators.required)
  });


  addSubscription(): void {

    let subName = this.addSubsriptionForm.get('name')?.getRawValue();
    let subPrice = this.addSubsriptionForm.get('price')?.getRawValue();

    if(subName != '' && subPrice != ''){

      this.api.createUserSubscription(subName, subPrice).subscribe({

        next: (res: Object) => {

          this.dialogRef.close({isCreated: true, name: subName});
        },
        error: (error: HttpErrorResponse) => {
          this.dialogs.errorDialog("Error Creating Subscription", "There was an error creating your subscription. Please try again later.");
        }

      });
    }    
  }

  cancel(): void {

    this.dialogRef.close({isCreated: false, name: ''});
  }

}
