import { HttpErrorResponse } from '@angular/common/http';
import { AfterViewInit, Component, EventEmitter, Input, OnChanges, Output, SimpleChanges, ViewChild } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatSort, Sort } from '@angular/material/sort';

@Component({
  selector: 'app-subscription-list',
  templateUrl: './subscription-list.component.html',
  styleUrls: ['./subscription-list.component.css']
})
export class SubscriptionListComponent implements AfterViewInit, OnChanges{

  constructor(private dialogs: DialogsService, private api: ApiService) {}

  @Input() subscriptionList: Subscription[] = [];
  @Output() updateSubscriptions = new EventEmitter<boolean>();   
  @Output() isActive = new EventEmitter<boolean>();
  @Output() isInactive = new EventEmitter<boolean>();
  @ViewChild(MatPaginator) paginator: MatPaginator = {} as MatPaginator;
  @ViewChild(MatSort) sort: MatSort = {} as MatSort;

  dataSource = new MatTableDataSource<Subscription>([]);
  active: boolean = true;
  currency: string = '$' ;

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  ngOnChanges(changes: SimpleChanges): void {    
    this.dataSource.data = changes['subscriptionList']["currentValue"];    
  }

  displayedColumns: string[] = ['name', 'price', 'dateadded', 'actions'];
  
  addSubscription(): void {
    this.dialogs.addSubscription().afterClosed().subscribe((res: {isCreated: boolean, name: string}) => {

      //successful creation of sub will fire this to add it
      if(res.isCreated === true){

        this.api.addUserSubscription(res.name).subscribe({

          next: (res: Object) => {
            this.updateSubscriptions.emit(true);
          },
          error: (error: HttpErrorResponse) => {

            if(error.status === 409){
              this.dialogs.errorDialog("Error Adding Subscription!", error["error"]["Error"]);
            }
            else{
              this.dialogs.errorDialog("Error Adding Subscription!", "There was an error in adding your subscription. Please try again later.")
            }
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
    
    this.displayedColumns = ['name', 'price', 'dateadded', 'actions'];
  }

  getInactive(): void {
    this.isInactive.emit(true);
    this.active = false;   
    
    this.displayedColumns = ['name', 'price', 'dateadded', 'dateremoved', 'actions'];
  }
}