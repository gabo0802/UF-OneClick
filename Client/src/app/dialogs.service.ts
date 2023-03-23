import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { ErrorComponent } from './dialogs/error/error.component';
import { SuccessComponent } from './dialogs/success/success.component';
import { PasswordResetComponent } from './dialogs/password-reset/password-reset.component';
import { DeleteAccountComponent } from './dialogs/delete-account/delete-account.component';
import { AddSubscriptionComponent } from './dialogs/add-subscription/add-subscription.component';
import { Subscription } from './subscription.model';
import { EditSubscriptionComponent } from './dialogs/edit-subscription/edit-subscription.component';

@Injectable({
  providedIn: 'root'
})
export class DialogsService {

  constructor(private dialog: MatDialog, private router: Router) { }


  successDialog(message: string): void {

    let dialogRef = this.dialog.open(SuccessComponent, {
      data: { dialogMessage: message},      
      height: '180px',
      width: '370px',
    });

    dialogRef.afterClosed().subscribe(result => {
        
      this.router.navigate(['/login']);      
    });

  }

  errorDialog(errorTitle: string, errorMessage: string): void {

    this.dialog.open(ErrorComponent, {
      data: {dialogTitle: errorTitle, dialogMessage: errorMessage},      
      height: '180px',
      width: '370px',
    });
  }

  deleteAccount(): void {

    this.dialog.open(DeleteAccountComponent, {           
      height: '200px',
      width: '410px',
    });
  }

  passwordReset(): void {

    this.dialog.open(PasswordResetComponent, { 
      height: '335px',
      width: '500px',
    }); 
  }

  addSubscription(): void {

    this.dialog.open(AddSubscriptionComponent, {
      height: '650px',
      width: '600px'
    });
  }

  editSubscription(subscriptionInfo: Subscription): void {

    this.dialog.open(EditSubscriptionComponent, {
      data: { subData: subscriptionInfo},
      height: '650px',
      width: '600px'
    });
  }
}
