package runtime

import "testing"

func TestInstallPackage(t *testing.T) {
	type args struct {
		remoteURL    string
		targetFolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InstallPackage(tt.args.remoteURL, tt.args.targetFolder); (err != nil) != tt.wantErr {
				t.Errorf("InstallPackage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
