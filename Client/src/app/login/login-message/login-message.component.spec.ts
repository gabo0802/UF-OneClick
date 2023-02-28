import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatDialogModule, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';
import { LoginMessageComponent } from './login-message.component';

describe('LoginMessageComponent', () => {
  let component: LoginMessageComponent;
  let fixture: ComponentFixture<LoginMessageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        LoginMessageComponent
       ],
       imports: [
        MaterialDesignModule,
        MatDialogModule
       ],
       providers: [
        { provide: MAT_DIALOG_DATA, useValue: {title: null, message: null}} 
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginMessageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
