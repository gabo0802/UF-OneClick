import { HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';

@Component({
  selector: 'app-reset',
  templateUrl: './reset.component.html',
  styleUrls: ['./reset.component.css']
})
export class ResetComponent {
  constructor(private api: ApiService, private dialogs: DialogsService) { }
  todayDate: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());

  adminReset(): void{
    if (confirm("Are you sure you want to reset?")){

      this.api.resetWebsite().subscribe({

        next: (res) => {
          this.dialogs.successDialog("Data successfully reset!");
        },
        error: (error: HttpErrorResponse) => {
          this.dialogs.errorDialog("Error Reseting Data!", "Error occurred resetting data!");
        }
      });
    }
  }
}
