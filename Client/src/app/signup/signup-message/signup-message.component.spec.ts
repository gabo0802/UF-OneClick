import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatDialog } from '@angular/material/dialog';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';

import { SignupMessageComponent } from './signup-message.component';

describe('SignupMessageComponent', () => {
  let component: SignupMessageComponent;
  let fixture: ComponentFixture<SignupMessageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        SignupMessageComponent,        
       ],
       imports: [
        MaterialDesignModule,
        HttpClientModule,                       
       ],
       providers: [
        MatDialog        
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SignupMessageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
