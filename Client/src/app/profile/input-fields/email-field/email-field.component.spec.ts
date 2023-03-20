import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';

import { EmailFieldComponent } from './email-field.component';

describe('EmailFieldComponent', () => {
  let component: EmailFieldComponent;
  let fixture: ComponentFixture<EmailFieldComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        EmailFieldComponent 
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

    fixture = TestBed.createComponent(EmailFieldComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('editing should be false initially', () => {
    expect(component.editing).toBeFalsy();
  });

  it('oldEmail should be empty before call', () => {
    expect(component.oldEmail).toEqual('');
  });

  it('Email form should be disabled initially', () => {
    expect(component.emailForm.disabled).toBeTruthy();
  });

  it('editing variable set to true when editEmail() initially called', () =>{

    component.editEmail();

    expect(component.editing).toBeTruthy();
  });

  it('when editEmail() called form is enabled', () => {

    component.editEmail()

    expect(component.emailForm.enabled).toBeTruthy();
  });

  it('when editEmail() called, form value is empty string', () => {

    component.editEmail()

    expect(component.emailForm.getRawValue()).toEqual('');
  });
});
