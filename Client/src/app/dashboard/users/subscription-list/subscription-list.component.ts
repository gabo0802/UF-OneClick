import { HttpErrorResponse } from '@angular/common/http';
import { AfterViewInit, Component, EventEmitter, Input, OnChanges, Output, SimpleChanges, ViewChild } from '@angular/core';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatSort } from '@angular/material/sort';
import { dateToString } from 'src/app/utils/dateToString';

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
  @ViewChild(MatSort) sort: MatSort = new MatSort;

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
  
  addActiveSubscription(): void {
    this.dialogs.addSubscription().afterClosed().subscribe((res: {isCreated: boolean, name: string, price: string, dateAdded: Date}) => {

      //successful creation of sub will fire this to add it
      if(res.isCreated === true){

        let todayDate: Date = new Date();
        
        //Checks if date added is equal to today's date
        if(this.dateChecker(todayDate, res.dateAdded)){

          this.api.addActiveUserSubscription(res.name).subscribe({

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
        else { //if active sub date added is before today's date

          let dateAddedString: string = dateToString(res.dateAdded);
          const subData = {name: res.name, price: res.price, dateadded: dateAddedString, dateremoved: ""};

          this.api.addOldUserSubscription(subData).subscribe({

            next: (res) => {  
              
              this.updateSubscriptions.emit(true);  
            },
            error: (error: HttpErrorResponse) => {
  
              if(error.status == 504){
                this.dialogs.errorDialog("Unexpected Error!", "An unexpected error occurred please try again later.");
              }
              else{
                this.dialogs.errorDialog("Error Adding Subscription", error["error"]["Error"]);
              }           
            }
          });

        }       
      }
    });
  }

  addInactiveSubscription(): void {
    this.dialogs.addInactiveSubscription();
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

  deleteSub(subID: string){
    
    this.api.deleteUserSubscription(subID).subscribe({

      next: (res) => {        
        this.getInactive();
      },
      error: (error: HttpErrorResponse) => {
        this.dialogs.errorDialog("Error Deleting Subscription!", "An error occured while trying to delete your subscription. Please try again later.");
      }
    });
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

  private dateChecker(date1: Date, date2: Date): boolean {

    let date1Year = date1.getFullYear();
    let date1Month = date1.getMonth();
    let date1Day = date1.getDate();

    let date2Year = date2.getFullYear();
    let date2Month = date2.getMonth();
    let date2Day = date2.getDate();

    return (date1Year === date2Year) && (date1Month === date2Month) && (date1Day === date2Day);
  }
}