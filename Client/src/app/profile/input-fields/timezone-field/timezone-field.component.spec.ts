import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';

import { TimezoneFieldComponent } from './timezone-field.component';

describe('TimezoneFieldComponent', () => {
  let component: TimezoneFieldComponent;
  let fixture: ComponentFixture<TimezoneFieldComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        TimezoneFieldComponent
       ],
       imports: [
        MaterialDesignModule, 
        HttpClientModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule
       ],

    })
    .compileComponents();

    fixture = TestBed.createComponent(TimezoneFieldComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('editing should be false initially', () => {
    expect(component.editing).toBeFalsy();
  });

  it('oldTimeZone should be empty before call', () => {
    expect(component.oldTimeZone).toEqual('');
  });

  it('Time Zone form should be disabled initially', () => {
    expect(component.timeZoneForm.disabled).toBeTruthy();
  });

  it('editing variable set to true when editTimeZone() initially called', () =>{

    component.editTimeZone();

    expect(component.editing).toBeTruthy();
  });

  it('when editTimeZone() called form is enabled', () => {

    component.editTimeZone()

    expect(component.timeZoneForm.enabled).toBeTruthy();
  });

  it('when editTimeZone() called form value is only "UTC"', () => {

    component.editTimeZone()

    expect(component.timeZoneForm.getRawValue()).toEqual('UTC');
  });

  it('when editing if editTimeZone() called, editing is false and form is disabled', () => {

    //enable editing
    component.editTimeZone()

    //disabled editing
    component.editTimeZone()

    expect(component.timeZoneForm.disabled).toBeTruthy();
    expect(component.editing).toBeFalsy();
  });

  it('Time Zone form has duplicate error if same Time Zone is entered', () => {

    //No error
    expect(component.timeZoneForm.hasError('duplicate')).toBeFalsy();

    component.oldTimeZone = '0500UTC';
    component.timeZoneForm.setValue('0500UTC');

    //triggers error
    component.updateTimeZone();

    expect(component.timeZoneForm.hasError('duplicate')).toBeTruthy();
  });
});
