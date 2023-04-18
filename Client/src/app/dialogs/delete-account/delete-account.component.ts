import { HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';

@Component({
  selector: 'app-delete-account',
  templateUrl: './delete-account.component.html',
  styleUrls: ['./delete-account.component.css']
})
export class DeleteAccountComponent {

  constructor(private api: ApiService, private dialogs: DialogsService, private router: Router) {}

  deleteAccount(): void {

    this.api.deleteUserAccount().subscribe({

      next: (res) => {
        this.dialogs.successDialog("Account succesfully deleted");
        this.router.navigate(['/']);
      },
      error: (error: HttpErrorResponse) => {
        this.dialogs.errorDialog("Error Deleting Accounting", "An Error occurred deleting your account. Please try again later.");
      }
    })
  }
}
