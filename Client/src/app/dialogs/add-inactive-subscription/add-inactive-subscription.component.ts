import { HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { AbstractControl, FormControl, FormGroup, ValidationErrors, ValidatorFn, Validators } from '@angular/forms';
import { MatOption } from '@angular/material/core';
import { MatDialogRef } from '@angular/material/dialog';
import { Observable, map, startWith, switchMap } from 'rxjs';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';
import { MatNativeDateModule } from '@angular/material/core';
import { MAT_DATEPICKER_VALIDATORS } from '@angular/material/datepicker';

@Component({
  selector: 'app-add-inactive-subscription',
  templateUrl: './add-inactive-subscription.component.html',
  styleUrls: ['./add-inactive-subscription.component.css']
})
export class AddInactiveSubscriptionComponent {

  constructor(private api: ApiService, private dialogs: DialogsService, private dialogRef: MatDialogRef<AddInactiveSubscriptionComponent>) {}

  allSubscriptions: Subscription[] = [];
  filteredOptions: Observable<Subscription[]> = new Observable<Subscription[]>;
  maxDate: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());
  minDate: Date = new Date(new Date().getFullYear()-40, new Date().getMonth(), new Date().getDate());

  addInactiveSubForm = new FormGroup({
    'name': new FormControl('', [Validators.required, Validators.pattern('^[ A-z0-9()\'+]+$')]),
    'price': new FormControl('', [Validators.required, Validators.pattern('^\\d{1,2}(.\\d\\d)?$')]),
    'dateadded': new FormControl<Date | null>(null,[Validators.required, this.datesValidator()]),
    'dateremoved': new FormControl<Date | null>(null,[Validators.required, this.datesValidator()])
  });

  ngOnInit(): void {       

    this.api.getAllSubscriptions().subscribe({

      next: (res: Subscription[]) => {

        this.allSubscriptions = res;

        //sorts subscriptions by name
        this.allSubscriptions.sort( (sub1, sub2) => {

          if(sub1.name < sub2.name){
            return -1;
          }

          if(sub1.name > sub2.name){
            return 1;
          }

          return 0;
        })

        this.populateDefaults();
      },
      error: (error: HttpErrorResponse) => {

        //initial message
        console.log(error.message);
      }
    });
  }

  private _filter(name: string): Subscription[] {
    const filterValue = name.toLowerCase();    

    return this.allSubscriptions.filter(subscription => subscription.name.toLowerCase().includes(filterValue));
  }

  displayFn(sub: Subscription): string {
    return sub && sub.name ? sub.name : '';
  }

  onSelected(option: MatOption) {  
    this.addInactiveSubForm.controls['price'].setValue(option.value.price);
    this.addInactiveSubForm.controls['price'].disable();    
  }

  private populateDefaults(): void {

    this.filteredOptions = this.addInactiveSubForm.controls['name'].valueChanges.pipe(
      startWith(''),
      map(value => {
        this.addInactiveSubForm.controls['price'].enable();        
        const name = typeof value === 'string' ? value : '';               
        return name ? this._filter(name as string) : this.allSubscriptions.slice();
      }),
    );

  }

  addSubscription(): void {

    
    let subName: string = '';
    let subPrice = this.addInactiveSubForm.get('price')?.getRawValue();
    let rawStartDate: Date = this.addInactiveSubForm.get('dateadded')?.getRawValue();
    let rawEndDate: Date = this.addInactiveSubForm.get('dateremoved')?.getRawValue();

    let subStartDate = this.dateToString(rawStartDate);
    let subEndDate = this.dateToString(rawEndDate);
    
    //Initial hack for checking whether it's a default sub
    //checks if not object
    if(typeof this.addInactiveSubForm.get('name')?.getRawValue() !== 'object'){

      subName = this.addInactiveSubForm.get('name')?.getRawValue();

      if(subName != '' && subPrice != '' && subStartDate !== null && subEndDate !== null){
      
        const subData = {name: subName, price: subPrice, dateadded: subStartDate, dateremoved: subEndDate};
        
        this.api.createUserSubscription(subName, subPrice).pipe(
          switchMap( res => {
            return this.api.addOldUserSubscription(subData);
          })
        ).subscribe({
          next: (res) => {
            console.log("New custom subscription added!");
            
          },
          error: (error: HttpErrorResponse) => {
            
            if(error.status == 504){
              this.dialogs.errorDialog("Unexpected Error!", "An unexpected error occurred please try again later.");
            }
            else{
              this.dialogs.errorDialog("Error Adding Inactive Subscription", error["error"]["Error"]);
            }
          }
        });

        this.dialogRef.close();
      }      
    }
    else {  
      
      //default sub should already be added;
      let subString = JSON.stringify(this.addInactiveSubForm.get('name')?.getRawValue());
      let subObject = JSON.parse(subString);
      
      subName = subObject.name;
      subPrice = subObject.price;

      if(subName != '' && subPrice != '' && subStartDate !== null && subEndDate !== null){

        const subData = {name: subName, price: subPrice, dateadded: subStartDate, dateremoved: subEndDate};

        this.api.addOldUserSubscription(subData).subscribe({

          next: (res) => {

            console.log("addedOld worked");

          },
          error: (error: HttpErrorResponse) => {

            if(error.status == 504){
              this.dialogs.errorDialog("Unexpected Error!", "An unexpected error occurred please try again later.");
            }
            else{
              this.dialogs.errorDialog("Error Adding Inactive Subscription", error["error"]["Error"]);
            }           
          }
        });
      }

      this.dialogRef.close();      
    }       
  }

  cancel(): void {

    this.dialogRef.close();
  }

  private dateToString(date: Date): string {

    //seconds between 10 and 60
    let randomSecs: number = Math.floor(Math.random() * 50 + 10);

    let month: string = (date.getMonth() + 1).toString();
    let dayDate: string = date.getDate().toString();    

    //if month is less than 10 adds leading zero
    if(date.getMonth() < 10){

      month = '0' + month;
    }

    //if date is less than 10 adds leading zero
    if(date.getDate() < 10) {
      
      dayDate = '0' + dayDate;
    }

    let stringDate = date.getFullYear() + "-" + month + "-" + dayDate + " 12:00:" + randomSecs;
    
    return stringDate;
  }

  //custom validator for checking date ranges
  //still has some errors, but works on a basic level
  private datesValidator(): ValidatorFn {

    return (group:AbstractControl) : ValidationErrors | null => {
  
      const startDate = group.parent?.get('dateadded')?.value;
      const endDate = group.parent?.get('dateremoved')?.value;      
      
      let status: boolean = true;

      if(startDate !== null && endDate !== null  && startDate < endDate){
        status = true;
      }
  
      if(startDate !== null && endDate !== null  && startDate > endDate){
        status = false;
      }      
  
      return status ? null : {range: true};
    }
  }

}