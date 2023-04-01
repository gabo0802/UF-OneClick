import { HttpErrorResponse } from '@angular/common/http';
import { AfterViewInit, Component, EventEmitter, Input, OnChanges, Output, SimpleChanges, ViewChild } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';

@Component({
  selector: 'app-subscription-list',
  templateUrl: './subscription-list.component.html',
  styleUrls: ['./subscription-list.component.css']
})
export class SubscriptionListComponent implements AfterViewInit, OnChanges{

  constructor(private dialogs: DialogsService, private api: ApiService) {}

  @Input() subscriptionList: Subscription[] = [];
  @Output() updateSubscriptions = new EventEmitter<boolean>();

  active: boolean = true;  
  @Output() isActive = new EventEmitter<boolean>();
  @Output() isInactive = new EventEmitter<boolean>();

  dataSource = new MatTableDataSource<Subscription>([]);

  @ViewChild(MatPaginator) paginator: MatPaginator = {} as MatPaginator;

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
  }

  ngOnChanges(changes: SimpleChanges): void {    
    this.dataSource.data = changes['subscriptionList']["currentValue"];
  }

  displayedColumns: string[] = ['sub-name', 'sub-price', 'sub-dateAdded', 'sub-actions'];
  
  addSubscription(): void {
    this.dialogs.addSubscription().afterClosed().subscribe((res: {isCreated: boolean, name: string}) => {

      //successful creation of sub will fire this to add it
      if(res.isCreated === true){

        this.api.addUserSubscription(res.name).subscribe({

          next: (res: Object) => {
            this.updateSubscriptions.emit(true);
          },
          error: (error: HttpErrorResponse) => {
            this.dialogs.errorDialog("Error Adding Subscription!", "There was an error in adding your subscription. Please try again later.")
          }
        });
      }
    });
  }

  deactivateSub(subName: string): void {
    
    this.api.deactivateSubscription(subName).subscribe({

      next: (res) => {
        
        this.updateSubscriptions.emit(true);
      },
      error: (error: HttpErrorResponse) => {
        this.dialogs.errorDialog("Error Deleting Subscription!", "There was an error deleting your subscription. Please try again later.");
      }
    })
  }

  reactivateSub(subName: string): void {

    this.api.reactivateSubscription(subName).subscribe({

      next: (res) => {

        //Potential other actions here.
        console.log("Reactivation success");
      },
      error: (error: HttpErrorResponse) => {

        this.dialogs.errorDialog("Error Adding Subscription", "Error adding!");
      }
    })
  }

  getActive(): void {
    this.isActive.emit(true);
    this.active = true;
    
    this.displayedColumns = ['sub-name', 'sub-price', 'sub-dateAdded', 'sub-actions'];
  }

  getInactive(): void {
    this.isInactive.emit(true);
    this.active = false;   
    
    this.displayedColumns = ['sub-name', 'sub-price', 'sub-dateAdded', 'sub-dateRemoved', 'sub-actions'];
  }
}