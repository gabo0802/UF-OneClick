import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ApiService } from '../api.service';
import { MaterialDesignModule } from '../material-design/material-design.module';
import { EmailFieldComponent } from './input-fields/email-field/email-field.component';
import { PasswordFieldComponent } from './input-fields/password-field/password-field.component';
import { TimezoneFieldComponent } from './input-fields/timezone-field/timezone-field.component';
import { UsernameFieldComponent } from './input-fields/username-field/username-field.component';

import { ProfileComponent } from './profile.component';

describe('ProfileComponent', () => {
  let component: ProfileComponent;
  let fixture: ComponentFixture<ProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        ProfileComponent,
        UsernameFieldComponent,
        EmailFieldComponent,
        TimezoneFieldComponent,
        PasswordFieldComponent,
       ],
       imports: [
        MaterialDesignModule, 
        HttpClientModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule
       ],
       providers: [
        ApiService
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
