# Changelog

이 프로젝트의 모든 주요 변경 사항은 이 파일에 기록됩니다.

## [v1.0.0] - 2025-12-31

### Added

- 입력 및 출력 디렉토리를 위한 CLI 플래그 추가 (`-input`, `-output`).
- Standard Go Project Layout 적용 (`cmd/`, `internal/`).
- 프로젝트 이력 추적을 위한 `CHANGELOG.md` 추가.
- MIT License 라이선스 파일(`LICENSE`) 추가.
- 필수 구성 파일 추가 (`Makefile`, `.editorconfig`, `.golangci.yml`, GitHub Actions CI).
- `.gitignore` 파일 확장 및 최적화.

### Changed

- `main.go` 로직을 `internal/organizer` 패키지로 리팩토링.
- 디렉토리 생성 권한을 `0600`에서 `0755`로 개선.
- `strings.Replacer`를 사용하여 문자열 정제 로직 최적화.
- 에러 처리 및 로깅 개선.

### Fixed

- 태그가 없을 때 변수 스코프 문제로 잘못된 경로가 생성되던 치명적인 버그 수정.
