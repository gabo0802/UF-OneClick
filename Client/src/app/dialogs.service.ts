import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { SuccessComponent } from './dialogs/success/success.component';

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
}
