import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';

import { AddSubscriptionComponent } from './add-subscription.component';

describe('AddSubscriptionComponent', () => {
  let component: AddSubscriptionComponent;
  let fixture: ComponentFixture<AddSubscriptionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        AddSubscriptionComponent 
      ],
      imports: [
        MaterialDesignModule, 
        HttpClientModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule,
       ],
       providers: [
        { provide: MatDialogRef<AddSubscriptionComponent>, useValue: { title: null} }
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddSubscriptionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('all Subscriptions should be initially empty', ()=> {

    expect(component.allSubscriptions).toEqual([]);
  });

  it('max Date should be todays date', () => {

    let todayDate = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());

    expect(component.maxDate).toEqual(todayDate);
  });

  it('minDate should be today date minus 40 years', () => {

    let minDate = new Date((new Date().getFullYear()) - 40, new Date().getMonth(), new Date().getDate());

    expect(component.minDate).toEqual(minDate);
  });

  it('form subscription name should initially be empty', ()=> {

    expect(component.addSubscriptionForm.get('name')?.getRawValue()).toEqual('');
  });

  it('form subscription price should initially be empty', ()=> {

    expect(component.addSubscriptionForm.get('price')?.getRawValue()).toEqual('');
  });

  it('form subscription date should initially be maxDate', ()=> {

    expect(component.addSubscriptionForm.get('dateadded')?.getRawValue()).toEqual(component.maxDate);
  });

  it('start date entered beyond maximimum limit throws maxDate error', () => {

    let startDate: Date = new Date(2025, 5, 15);

    component.addSubscriptionForm.get('dateadded')?.setValue(startDate);

    expect(component.addSubscriptionForm.get('dateadded')?.hasError('matDatepickerMax')).toEqual(true);
  });

  it('start date entered below minimum limit throws minDate error', ()=> {

    let startDate: Date = new Date(1972, 5, 15);

    component.addSubscriptionForm.get('dateadded')?.setValue(startDate);

    expect(component.addSubscriptionForm.get('dateadded')?.hasError('matDatepickerMin')).toEqual(true);
  });
});
