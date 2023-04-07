import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddInactiveSubscriptionComponent } from './add-inactive-subscription.component';
import { HttpClientModule } from '@angular/common/http';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatDialogRef } from '@angular/material/dialog';

describe('AddInactiveSubscriptionComponent', () => {
  let component: AddInactiveSubscriptionComponent;
  let fixture: ComponentFixture<AddInactiveSubscriptionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        AddInactiveSubscriptionComponent
      ],
      imports: [
        HttpClientModule, 
        MaterialDesignModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule
       ],
       providers: [
        { provide: MatDialogRef, useValue: {title: null}} 
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddInactiveSubscriptionComponent);
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

    expect(component.addInactiveSubForm.get('name')?.getRawValue()).toEqual('');
  });

  it('form subscription price should initially be empty', ()=> {

    expect(component.addInactiveSubForm.get('price')?.getRawValue()).toEqual('');
  });

  it('form subscription dateadded should initially be null', ()=> {

    expect(component.addInactiveSubForm.get('dateadded')?.getRawValue()).toEqual(null);
  });

  it('form subscription dateremoved should initially be null', ()=> {

    expect(component.addInactiveSubForm.get('dateremoved')?.getRawValue()).toEqual(null);
  });

  it('Dates Validator function checking if error thrown on end date form if start date after end date', ()=> {

    let startDate: Date = new Date(2011, 4, 25);
    let endDate: Date = new Date(2010, 5, 15);

    component.addInactiveSubForm.get('dateadded')?.setValue(startDate);
    component.addInactiveSubForm.get('dateremoved')?.setValue(endDate);

    expect(component.addInactiveSubForm.get('dateremoved')?.hasError('range')).toEqual(true);
  });

  it('Dates Validator function checking if error thrown on start end if start after end date', ()=> {

    let startDate: Date = new Date(2012, 6, 25);
    let endDate: Date = new Date(2008, 5, 15);

    component.addInactiveSubForm.get('dateremoved')?.setValue(endDate);
    component.addInactiveSubForm.get('dateadded')?.setValue(startDate);    

    expect(component.addInactiveSubForm.get('dateadded')?.hasError('range')).toEqual(true);
  });

  it('start date entered beyond maximimum limit throws maxDate error', () => {

    let startDate: Date = new Date(2025, 5, 15);

    component.addInactiveSubForm.get('dateadded')?.setValue(startDate);

    expect(component.addInactiveSubForm.get('dateadded')?.hasError('matDatepickerMax')).toEqual(true);
  });

  it('start date entered below minimum limit throws minDate error', ()=> {

    let startDate: Date = new Date(1972, 5, 15);

    component.addInactiveSubForm.get('dateadded')?.setValue(startDate);

    expect(component.addInactiveSubForm.get('dateadded')?.hasError('matDatepickerMin')).toEqual(true);
  });

  it('end date entered beyond maximimum limit throws maxDate error', () => {

    let endDate: Date = new Date(2025, 5, 15);

    component.addInactiveSubForm.get('dateremoved')?.setValue(endDate);

    expect(component.addInactiveSubForm.get('dateremoved')?.hasError('matDatepickerMax')).toEqual(true);
  });

  it('end date entered below minimum limit throws minDate error', ()=> {

    let endDate: Date = new Date(1972, 5, 15);

    component.addInactiveSubForm.get('dateremoved')?.setValue(endDate);

    expect(component.addInactiveSubForm.get('dateremoved')?.hasError('matDatepickerMin')).toEqual(true);
  });
});
