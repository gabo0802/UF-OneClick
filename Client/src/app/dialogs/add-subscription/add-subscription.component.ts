import { HttpErrorResponse } from '@angular/common/http';
import { AfterViewInit, Component, ComponentFactoryResolver, OnChanges, OnInit, SimpleChanges } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatOption } from '@angular/material/core';
import { MatDialogRef } from '@angular/material/dialog';
import { Observable, map, startWith} from 'rxjs';
import { ApiService } from 'src/app/api.service';
import { DialogsService } from 'src/app/dialogs.service';
import { Subscription } from 'src/app/subscription.model';

@Component({
  selector: 'app-add-subscription',
  templateUrl: './add-subscription.component.html',
  styleUrls: ['./add-subscription.component.css']
})
export class AddSubscriptionComponent implements OnInit{

  constructor(private api: ApiService, private dialogs: DialogsService, private dialogRef: MatDialogRef<AddSubscriptionComponent>) {}

  allSubscriptions: Subscription[] = [];
  filteredOptions: Observable<Subscription[]> = new Observable<Subscription[]>;
  
  addSubscriptionForm = new FormGroup({    
    'name': new FormControl('', [Validators.required, Validators.pattern('^[ A-z0-9()\'+]+$')]),
    'price': new FormControl('', [Validators.required, Validators.pattern('^\\d{1,2}(.\\d\\d)?$')])
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
    this.addSubscriptionForm.controls['price'].setValue(option.value.price);
    this.addSubscriptionForm.controls['price'].disable();    
  }

  populateDefaults(): void {

    this.filteredOptions = this.addSubscriptionForm.controls['name'].valueChanges.pipe(
      startWith(''),
      map(value => {
        this.addSubscriptionForm.controls['price'].enable();        
        const name = typeof value === 'string' ? value : '';               
        return name ? this._filter(name as string) : this.allSubscriptions.slice();
      }),
    );

  }

  addSubscription(): void {

    let subName: string = '';
    let subPrice = this.addSubscriptionForm.get('price')?.getRawValue();
    
    //Initial hack for checking whether it's a default sub
    //checks if not object
    if(typeof this.addSubscriptionForm.get('name')?.getRawValue() !== 'object'){

      subName = this.addSubscriptionForm.get('name')?.getRawValue();

      if(subName != '' && subPrice != ''){
      
        this.api.createUserSubscription(subName, subPrice).subscribe({
  
          next: (res) => {
            
            this.dialogRef.close({isCreated: true, name: subName});
          },
          error: (error: HttpErrorResponse) => {

            if(error.status == 409){
              this.addSubscriptionForm.controls['name'].setErrors({'duplicate': true});
            }
            
            this.dialogs.errorDialog("Error Creating Subscription", error["error"]["Error"]);
          }
  
        });
      }
      
    }
    else {
      
      //If default sub, then already valid and sends back success

      let subString = JSON.stringify(this.addSubscriptionForm.get('name')?.getRawValue());
      let subObject = JSON.parse(subString);
      subName = subObject.name;

      this.dialogRef.close({isCreated: true, name: subName});      
    }       
  }

  cancel(): void {

    this.dialogRef.close({isCreated: false, name: ''});
  }

}
